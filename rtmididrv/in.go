package rtmididrv

import (
	"fmt"
	"math"

	"github.com/gomidi/connect"
	"github.com/gomidi/connect/imported/rtmidi"
)

type in struct {
	driver *driver
	number int
	name   string
	isOpen bool
}

// IsOpen returns wether the MIDI in port is open
func (i *in) IsOpen() bool {
	return i.isOpen
}

// String returns the name of the MIDI in port.
func (i *in) String() string {
	return i.name
}

// Underlying returns the underlying rtmidi.MIDIIn. Use it with type casting:
//   rtIn := i.Underlying().(rtmidi.MIDIIn)
func (i *in) Underlying() interface{} {
	return i.driver.in
}

// Number returns the number of the MIDI in port.
// Note that with rtmidi, out and in ports are counted separately.
// That means there might exists out ports and an in ports that share the same number.
func (i *in) Number() int {
	return i.number
}

// Close closes the MIDI in port, after it has stopped listening.
func (i *in) Close() error {
	if !i.isOpen {
		return nil
	}

	err := i.StopListening()
	if err != nil {
		return err
	}

	err = i.driver.in.Close()
	if err != nil {
		return fmt.Errorf("can't close MIDI in port %v (%s): %v", i.number, i, err)
	}

	i.isOpen = false
	return nil
}

// Open opens the MIDI in port
func (i *in) Open() error {
	err := i.driver.in.OpenPort(i.number, "")
	if err != nil {
		return fmt.Errorf("can't open MIDI in port %v (%s): %v", i.number, i, err)
	}
	i.isOpen = true
	return nil
}

func newIn(a *driver, number int, name string) connect.In {
	return &in{driver: a, number: number, name: name, isOpen: false}
}

// SetListener makes the listener listen to the in port
func (i *in) SetListener(listener func(data []byte, deltaMicroseconds int64)) error {
	err := i.driver.in.SetCallback(func(_ rtmidi.MIDIIn, bt []byte, deltaSeconds float64) {
		// we want deltaMicroseconds as int64
		listener(bt, int64(math.Round(deltaSeconds*1000000)))
	})
	if err != nil {
		fmt.Errorf("can't set listener for MIDI in port %v (%s): %v", i.number, i, err)
	}
	return nil
}

// StopListening cancels the listening
func (i *in) StopListening() error {
	err := i.driver.in.CancelCallback()
	if err != nil {
		fmt.Errorf("can't stop listening on MIDI in port %v (%s): %v", i.number, i, err)
	}
	return nil
}
