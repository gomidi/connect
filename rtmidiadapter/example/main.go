package main

import (
	"fmt"
	"github.com/gomidi/connect/imported/rtmidi"
	rta "github.com/gomidi/connect/rtmidiadapter"
	"github.com/gomidi/mid"
	"time"
)

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {

	{ // find the ports
		printInPorts()
		printOutPorts()
		fmt.Println(" \n--Messages--")
	}

	var ( // wire it up
		rtOut   = openOut(0) // where we write to, customize the port!
		rtIn    = openIn(0)  // where we listen on, customize the port!
		in, out = rta.In(rtIn), rta.Out(rtOut)
		rd      = mid.NewReader()
		wr      = mid.WriteTo(out)
	)

	// listen for MIDI
	go rd.ReadFrom(in)

	{ // write MIDI to out that passes it to in on which we listen.
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

		// close the rtmidi ports
		rtIn.Destroy()
		rtOut.Destroy()
	}
}

func openIn(port int) rtmidi.MIDIIn {
	in, err := rta.OpenIn(port)
	if err != nil {
		panic(err.Error())
	}
	return in
}

func openOut(port int) rtmidi.MIDIOut {
	out, err := rta.OpenOut(port)
	if err != nil {
		panic(err.Error())
	}
	return out
}

func printInPorts() {
	ins, err := rta.InPorts()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\n---MIDI input ports---")

	for i, name := range ins {
		fmt.Printf("%d %#v\n", i, name)
	}
}

func printOutPorts() {
	outs, err := rta.OutPorts()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\n---MIDI output ports---")

	for i, name := range outs {
		fmt.Printf("%d %#v\n", i, name)
	}
}
