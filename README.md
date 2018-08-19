# connect
Adapters to connect with the outside MIDI world

## Purpose

To connect the https://github.com/gomidi/mid package with the outside MIDI world, there are two adapter packages - one for `rtmidi` and one for `portmidi`.

## Installation

It is recommended to use Go 1.11 with module support (`$GO111MODULE=on`).

For rtmidi (see https://github.com/thestk/rtmidi for more information)

```
go get -d github.com/gomidi/connect/rtmidiadapter
```

For portaudio (see https://github.com/rakyll/portmidi for more information)

```
// install the headers of portmidi somehow, e.g. apt-get install libportmidi-dev
go get -d github.com/gomidi/connect/portmidiadapter
```

## Example

This example uses `rtmidi` which in my experience is far more perfomant than `portmidi`.

An example with portmidi could be found at `portmidiadapter/example`.

```go
package main

import (
    "fmt"
    "github.com/gomidi/connect/imported/rtmidi"
    rta "github.com/gomidi/connect/rtmidiadapter"
    "github.com/gomidi/mid"
    "time"
)

func main() {

    { // find the ports
        printInPorts()
        printOutPorts()
        fmt.Println(" \n--Messages--")
    }

    var ( // wire it up
        rtIn, rtOut = openIn(0), openOut(0)
        in, out     = rta.In(rtIn), rta.Out(rtOut)
        rd          = mid.NewReader()
        wr          = mid.SpeakTo(out)
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
```
