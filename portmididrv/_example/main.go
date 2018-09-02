package main

import (
	"time"

	"github.com/gomidi/connect/portmididrv"
	"github.com/gomidi/mid"
)

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {

	// don't forget!
	drv, err := portmididrv.New()

	if err != nil {
		panic(err.Error())
	}

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	if err != nil {
		panic("can't find MIDI in ports")
	}

	outs, err := drv.Outs()
	if err != nil {
		panic("can't find MIDI out ports")
	}

	rd := mid.NewReader()
	wr := mid.WriteTo(outs[0])

	// listen for MIDI
	go rd.ReadFrom(ins[0])

	{ // write MIDI
		wr.NoteOn(60, 100)
		time.Sleep(time.Nanosecond)
		wr.NoteOff(60)
		time.Sleep(time.Nanosecond)

		wr.SetChannel(1)

		wr.NoteOn(70, 100)
		time.Sleep(time.Nanosecond)
		wr.NoteOff(70)
		time.Sleep(time.Second * 1)
	}

	// close the portmidi ports (would be done via drv.Close() anyway
	ins[0].Close()
	outs[0].Close()
}
