# connect
Go interface for MIDI drivers


## Purpose

Unification of MIDI driver packages for Go. Currently two implementations exist: 
- for rtmidi: https://github.com/gomidi/rtmididrv
- for portmidi: https://github.com/gomidi/portmididrv

This package is also used by https://github.com/gomidi/mid for smooth integration

## Installation

It is recommended to use Go 1.11 with module support (`$GO111MODULE=on`).

For rtmidi (see https://github.com/thestk/rtmidi for more information)

```
// install the headers of alsa somehow, e.g. sudo apt-get install libasound2-dev
go get -d github.com/gomidi/rtmididrv
```

For portaudio (see https://github.com/rakyll/portmidi for more information)

```
// install the headers of portmidi somehow, e.g. sudo apt-get install libportmidi-dev
go get -d github.com/gomidi/portmididrv
```

## Documentation

rtmididrv: [![rtmidi docs](http://godoc.org/github.com/gomidi/rtmididrv?status.png)](http://godoc.org/github.com/gomidi/rtmididrv)

portmididrv: [![portmidi docs](http://godoc.org/github.com/gomidi/portmididrv?status.png)](http://godoc.org/github.com/gomidi/portmididrv)

## Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomidi/connect"
	"github.com/gomidi/mid"
	driver "github.com/gomidi/rtmididrv"
	// for portmidi
	// driver "github.com/gomidi/portmididrv" 
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {
	drv, err := driver.New()
	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	if len(os.Args) == 2 && os.Args[1] == "list" {
		printInPorts(ins)
		printOutPorts(outs)
		return
	}

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	wr := mid.WriteTo(out)

	// listen for MIDI
	go mid.NewReader().ReadFrom(in)

	{ // write MIDI to out that passes it to in on which we listen.
		err := wr.NoteOn(60, 100)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Nanosecond)
		wr.NoteOff(60)
		time.Sleep(time.Nanosecond)

		wr.SetChannel(1)

		wr.NoteOn(70, 100)
		time.Sleep(time.Nanosecond)
		wr.NoteOff(70)
		time.Sleep(time.Second * 1)
	}
}

func printPort(port connect.Port) {
	fmt.Printf("[%v] %s\n", port.Number(), port.String())
}

func printInPorts(ports []connect.In) {
	fmt.Printf("MIDI IN Ports\n")
	for _, port := range ports {
		printPort(port)
	}
	fmt.Printf("\n\n")
}

func printOutPorts(ports []connect.Out) {
	fmt.Printf("MIDI OUT Ports\n")
	for _, port := range ports {
		printPort(port)
	}
	fmt.Printf("\n\n")
}

```

## Example 2 reading from device  

```go
package main

import (
	"fmt"

	"github.com/gomidi/connect"
	driver "github.com/gomidi/rtmididrv"
)

func main() {
	// Create a new midi driver and close it with "defer" when the program terminates
	var drv connect.Driver
	drv, _ = driver.New()

	//get the input devices
	fmt.Println(getMIDIDevices(drv))

	// open the device in index "0", without using device id
	in, _ := connect.OpenIn(drv, 0, "")

	// open the device by passing the name, and giving an id >= 0
	//example connect.OpenIn(drv, 1, "AKAI 0")

	//in, _ := connect.OpenIn(drv, 1, getMIDIDevices(drv)[0].String())

	in.SetListener(handleMIDIevent)

	// we can keep the thread running by keeping the while loop
	for {

	}
}

// gets all the midi devices from the driver and returns them
func getMIDIDevices(drv connect.Driver) (outs []connect.In) {
	//gets the inputs
	ins, _ := drv.Ins()

	/* // that is the way to get the outputs
	ins, _ := drv.Outs() */

	for _, el := range ins {
		outs = append(outs, el)
	}
	return outs
}

func handleMIDIevent(data []byte, deltaMicroseconds int64) {
	fmt.Printf("Midi message on: %v value: %v \n", data[1], data[2])
}

```  
```  
example output would be:  
[LPD8 0]
Midi message on: 6 value: 4
Midi message on: 6 value: 127
Midi message on: 2 value: 66
Midi message on: 2 value: 127
Midi message on: 7 value: 4
Midi message on: 7 value: 127
```  
