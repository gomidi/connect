package rtmididrv

import (
	"fmt"

	"github.com/gomidi/connect"
)

func newOut(d *driver, number int, name string) connect.Out {
	return &out{driver: d, number: number, name: name, isOpen: false}
}

type out struct {
	driver *driver
	number int
	name   string
	isOpen bool
}

// IsOpen returns wether the port is open
func (o *out) IsOpen() bool {
	return o.isOpen
}

// Send sends a message to the MIDI out port
// If the out port is closed, it returns connect.ErrClosed
func (o *out) Send(b []byte) error {
	if !o.isOpen {
		return connect.ErrClosed
	}
	err := o.driver.out.SendMessage(b)
	if err != nil {
		return fmt.Errorf("could not send message to MIDI out %v (%s): %v", o.number, o, err)
	}
	return nil
}

// Underlying returns the underlying rtmidi.MIDIOut. Use it with type casting:
//   rtOut := o.Underlying().(rtmidi.MIDIOut)
func (o *out) Underlying() interface{} {
	return o.driver.out
}

// Number returns the number of the MIDI out port.
// Note that with rtmidi, out and in ports are counted separately.
// That means there might exists out ports and an in ports that share the same number
func (o *out) Number() int {
	return o.number
}

// String returns the name of the MIDI out port.
func (o *out) String() string {
	return o.name
}

// Close closes the MIDI out port
func (o *out) Close() error {
	if !o.isOpen {
		return nil
	}
	err := o.driver.out.Close()
	if err != nil {
		return fmt.Errorf("can't close MIDI out %v (%s): %v", o.number, o, err)
	}
	o.isOpen = false
	return nil
}

// Open opens the MIDI out port
func (o *out) Open() error {
	err := o.driver.out.OpenPort(o.number, "")
	if err != nil {
		return fmt.Errorf("can't open MIDI out port %v (%s): %v", o.number, o, err)
	}
	o.isOpen = true
	return nil
}
