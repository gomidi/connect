package main

import (
	"bytes"
	"fmt"
	"github.com/gomidi/connect/portmidiadapter"
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
		printPorts()
		fmt.Println(" \n--Messages--")
	}

	var ( // wire it up
		midiOut = openMIDIOut(portmidi.DefaultOutputDeviceID()) // where we write to, customize the port!
		midiIn  = openMIDIIn(portmidi.DefaultInputDeviceID())   // where we listen on, customize the port!
		in, out = portmidiadapter.In(midiIn), portmidiadapter.Out(midiOut)
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
		midiIn.Close()
		midiOut.Close()
	}
}

func openMIDIIn(port portmidi.DeviceID) *portmidi.Stream {
	in, err := portmidi.NewInputStream(port, 1024)

	if err != nil {
		panic("can't open MIDI in port:" + err.Error())
	}

	return in
}

func openMIDIOut(port portmidi.DeviceID) *portmidi.Stream {
	out, err := portmidi.NewOutputStream(port, 1024, 0)

	if err != nil {
		panic("can't open MIDI out port:" + err.Error())
	}

	return out
}

func printPorts() {
	var ins, outs bytes.Buffer

	no := portmidi.CountDevices()

	for i := 0; i < no; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		if info.IsInputAvailable {
			fmt.Fprintf(&ins, "%d %#v\n", i, info.Name)
		}

		if info.IsOutputAvailable {
			fmt.Fprintf(&outs, "%d %#v\n", i, info.Name)
		}
	}

	fmt.Println("\n---MIDI input ports---")
	fmt.Println(ins.String())

	fmt.Println("\n---MIDI output ports---")
	fmt.Println(outs.String())

}
