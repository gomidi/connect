package main

import (
	"fmt"
	pma "github.com/gomidi/connect/portmidiadapter"
	"github.com/gomidi/mid"
	"github.com/rakyll/portmidi"
	"time"
)

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {

	// don't forget!
	portmidi.Initialize()

	{ // find the ports
		printInPorts()
		printOutPorts()
		fmt.Println(" \n--Messages--")
	}

	var ( // wire it up
		pmOut   = openOut(1) // where we write to, customize the port!
		pmIn    = openIn(0)  // where we listen on, customize the port!
		in, out = pma.In(pmIn), pma.Out(pmOut)
		rd      = mid.NewReader()
		wr      = mid.SpeakTo(out)
	)

	// listen for MIDI
	go rd.ListenTo(in)

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

	{ // clean up
		in.StopListening()
		pmIn.Close()
		pmOut.Close()
	}
}

func openIn(port portmidi.DeviceID) *portmidi.Stream {
	in, err := pma.OpenIn(port)
	if err != nil {
		panic("can't open MIDI in port:" + err.Error())
	}
	return in
}

func openOut(port portmidi.DeviceID) *portmidi.Stream {
	out, err := pma.OpenOut(port)
	if err != nil {
		panic("can't open MIDI out port:" + err.Error())
	}
	return out
}

func printInPorts() {
	ins := pma.InPorts()
	fmt.Println("\n---MIDI input ports---")
	for i, name := range ins {
		fmt.Printf("%d %#v\n", i, name)
	}
}

func printOutPorts() {
	outs := pma.OutPorts()
	fmt.Println("\n---MIDI output ports---")
	for i, name := range outs {
		fmt.Printf("%d %#v\n", i, name)
	}
}
