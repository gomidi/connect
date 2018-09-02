package rtmididrv

import (
	"fmt"

	"github.com/gomidi/connect"
	"github.com/gomidi/connect/imported/rtmidi"
)

type driver struct {
	in  rtmidi.MIDIIn
	out rtmidi.MIDIOut
}

func (d *driver) String() string {
	return "rtmididrv"
}

func (d *driver) Close() error {
	d.in.Destroy()
	d.out.Destroy()
	return nil
}

// Custom returns a driver based on a custom rtmidi.MIDIIn and rtmidi.MIDIOut
func Custom(in rtmidi.MIDIIn, out rtmidi.MIDIOut) connect.Driver {
	return &driver{in, out}
}

// New returns a driver based on the default rtmidi in and out
func New() (connect.Driver, error) {
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI in: %v", err)
	}
	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI out: %v", err)
	}
	return &driver{in, out}, nil
}

// Ins returns the available MIDI input ports
func (a *driver) Ins() (ins []connect.In, err error) {
	ports, err := a.in.PortCount()
	if err != nil {
		return nil, fmt.Errorf("can't get number of in ports: %s", err.Error())
	}

	for i := 0; i < ports; i++ {
		name, err := a.in.PortName(i)
		if err != nil {
			name = ""
		}
		ins = append(ins, newIn(a, i, name))
	}
	return
}

// Outs returns the available MIDI output ports
func (a *driver) Outs() (outs []connect.Out, err error) {
	ports, err := a.out.PortCount()
	if err != nil {
		return nil, fmt.Errorf("can't get number of out ports: %s", err.Error())
	}

	for i := 0; i < ports; i++ {
		name, err := a.out.PortName(i)
		if err != nil {
			name = ""
		}
		outs = append(outs, newOut(a, i, name))
	}
	return
}
