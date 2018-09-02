# connect
Go interface for MIDI drivers

[![Build Status Travis/Linux](https://travis-ci.org/gomidi/connect.svg?branch=master)](http://travis-ci.org/gomidi/connect)

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
	"time"

	"github.com/gomidi/rtmididrv"
	// for portmidi
	// "github.com/gomidi/portmididrv"
	"github.com/gomidi/mid"
)

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {
	drv := rtmididrv.New()
	
	// for portmidi
    // drv, err := portrtmididrv.New()

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

	// close the ports (would be done via drv.Close() anyway
	ins[0].Close()
	outs[0].Close()
}

```
