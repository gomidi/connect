# connect
Adapters to connect with the outside MIDI world

## Purpose

To connect the https://github.com/gomidi/mid package with the outside MIDI world, there are two adapter packages.

## Installation

It is recommended to use Go 1.11 with module support (`$GO111MODULE=on`).

For rtmidi (see https://github.com/thestk/rtmidi for more information)

```
// install the headers of rtmidi somehow, e.g. apt-get install librtmidi2v5
go get -d github.com/gomidi/connect/rtmidiadapter
```

or

For portaudio (see https://github.com/rakyll/portmidi for more information)

```
// install the headers of portmidi somehow, e.g. apt-get install libportmidi0 libportmidi-dev
go get -d github.com/gomidi/connect/portmidiadapter
```

## Example with rtmidi

```go
package main

import (
    "fmt"
    "github.com/gomidi/connect/imported/rtmidi"
    "github.com/gomidi/connect/rtmidiadapter"
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
        midiIn, midiOut = openMIDIIn(0), openMIDIOut(0)
        in, out         = rtmidiadapter.In(midiIn), rtmidiadapter.Out(midiOut)
        rd              = mid.NewReader()
        wr              = mid.SpeakTo(out)
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
        midiIn.Destroy()
        midiOut.Destroy()
    }
}

func openMIDIIn(port int) rtmidi.MIDIIn {
    in, err := rtmidi.NewMIDIInDefault()
    if err != nil {
        panic("can't open default MIDI in:" + err.Error())
    }

    err = in.OpenPort(port, "")
    if err != nil {
        panic("can't open MIDI in port:" + err.Error())
    }

    return in
}

func openMIDIOut(port int) rtmidi.MIDIOut {
    out, err := rtmidi.NewMIDIOutDefault()
    if err != nil {
        panic("can't open default MIDI out:" + err.Error())
    }

    err = out.OpenPort(port, "")
    if err != nil {
        panic("can't open MIDI out port:" + err.Error())
    }

    return out
}

func printInPorts() {
    in, err := rtmidi.NewMIDIInDefault()
    if err != nil {
        panic("can't open default MIDI in:" + err.Error())
    }

    ports, errP := in.PortCount()
    if errP != nil {
        panic("can't get number of in ports:" + errP.Error())
    }

    fmt.Println("\n---MIDI input ports---")

    for i := 0; i < ports; i++ {
        name, _ := in.PortName(i)
        fmt.Printf("%d %#v\n", i, name)
    }
}

func printOutPorts() {
    out, err := rtmidi.NewMIDIOutDefault()
    if err != nil {
        panic("can't open default MIDI out:" + err.Error())
    }

    ports, errP := out.PortCount()
    if errP != nil {
        panic("can't get number of out ports:" + errP.Error())
    }

    fmt.Println("\n---MIDI output ports---")

    for i := 0; i < ports; i++ {
        name, _ := out.PortName(i)
        fmt.Printf("%d %#v\n", i, name)
    }
}

```

## Example with portmidi

```go
package main

import (
    "bytes"
    "fmt"
    "github.com/gomidi/connect/portmidiadapter"
    "github.com/gomidi/mid"
    "github.com/rakyll/portmidi"
    "time"
)

func main() {

    // don't forget!
    portmidi.Initialize()

    { // find the ports
        printPorts()
        fmt.Println(" \n--Messages--")
    }

    var ( // wire it up
        midiIn  = openMIDIIn(portmidi.DefaultInputDeviceID())
        midiOut = openMIDIOut(portmidi.DefaultOutputDeviceID())
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

```