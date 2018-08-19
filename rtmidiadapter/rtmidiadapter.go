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

		// fmt.Printf("%d %#v\n", i, name)
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

	/*
		// fmt.Println("\n---MIDI input ports---")
		var names []string

		for i := 0; i < ports; i++ {
			name, errN := in.PortName(i)

			if errN != nil {
				name = ""
			}

			names = append(names, name)

			// fmt.Printf("%d %#v\n", i, name)
		}
	*/

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

	// fmt.Println("\n---MIDI output ports---")

	/*
		for i := 0; i < ports; i++ {
			name, _ := out.PortName(i)
			fmt.Printf("%d %#v\n", i, name)
		}
	*/
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

func Out(o rtmidi.MIDIOut) mid.Out {
	return &out{o}
}

type out struct {
	rtmidi.MIDIOut
}

func (o *out) Send(b []byte) error {
	return o.SendMessage(b)
}

func In(i rtmidi.MIDIIn) mid.In {
	return &in{i}
}

type in struct {
	rtmidi.MIDIIn
}

func (i *in) StopListening() {
	i.CancelCallback()
}

func (i *in) SetListener(f func([]byte)) {
	i.SetCallback(func(_ rtmidi.MIDIIn, bt []byte, _ float64) {
		f(bt)
	})
}
