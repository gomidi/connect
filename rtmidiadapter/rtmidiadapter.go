package rtmidiadapter

import (
	"fmt"
	"github.com/gomidi/mid"
	//github.com/thestk/rtmidi/tree/master/contrib/go/rtmidi
	"github.com/gomidi/connect/imported/rtmidi"
)

func names(m rtmidi.MIDI, numports int) (n []string) {
	for i := 0; i < numports; i++ {
		name, err := m.PortName(i)

		if err != nil {
			name = ""
		}

		n = append(n, name)
	}

	return
}

// InPorts returns a slice of names of the input ports.
// The corresponding port number is the corresponding index.
// If a name got not be retrieved, the string is empty.
func InPorts() ([]string, error) {
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI in: %s", err.Error())
	}

	ports, errP := in.PortCount()
	if errP != nil {
		return nil, fmt.Errorf("can't get number of in ports: %s", errP.Error())
	}

	return names(in, ports), nil
}

// OutPorts returns a slice of names of the output ports.
// The corresponding port number is the corresponding index.
// If a name got not be retrieved, the string is empty.
func OutPorts() ([]string, error) {
	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI out: %s", err.Error())
	}

	ports, errP := out.PortCount()
	if errP != nil {
		return nil, fmt.Errorf("can't get number of out ports: %s", errP.Error())
	}

	return names(out, ports), nil
}

// OpenIn opens the given input port and returns it.
func OpenIn(port int) (rtmidi.MIDIIn, error) {
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI in")
	}

	err = in.OpenPort(port, "")
	if err != nil {
		return nil, fmt.Errorf("can't open MIDI in port %v", port)
	}

	return in, nil
}

// OpenOut opens the given output port and returns it.
func OpenOut(port int) (rtmidi.MIDIOut, error) {
	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		return nil, fmt.Errorf("can't open default MIDI out")
	}

	err = out.OpenPort(port, "")
	if err != nil {
		return nil, fmt.Errorf("can't open MIDI out port %v", port)
	}

	return out, nil
}

// Out wraps an rtmidi.MIDIOut to make it compatible with gomidi/mid#Out
func Out(o rtmidi.MIDIOut) mid.Out {
	return &out{o}
}

type out struct {
	rtmidi.MIDIOut
}

// Send sends a message to the out port
func (o *out) Send(b []byte) error {
	return o.SendMessage(b)
}

// In wraps an rtmidi.MIDIIn to make it compatible with gomidi/mid#In
func In(i rtmidi.MIDIIn) mid.In {
	return &in{i}
}

type in struct {
	rtmidi.MIDIIn
}

// StopListening cancels the listening
func (i *in) StopListening() {
	i.CancelCallback()
}

// SetListener makes the listener listen to the in port
func (i *in) SetListener(f func([]byte)) {
	i.SetCallback(func(_ rtmidi.MIDIIn, bt []byte, _ float64) {
		f(bt)
	})
}
