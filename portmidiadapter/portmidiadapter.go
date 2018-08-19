package portmidiadapter

import (
	"github.com/gomidi/mid"
	"github.com/rakyll/portmidi"
	"sync"
	"time"
)

// Out wraps an output portmidi.Stream to make it compatible with gomidi/mid#Out
func Out(o *portmidi.Stream) mid.Out {
	return &out{o}
}

type out struct {
	*portmidi.Stream
}

func (o *out) Send(b []byte) error {
	return o.WriteShort(int64(b[0]), int64(b[1]), int64(b[2]))
}

// InOption is an option that can be passed to In
type InOption func(*in)

// BufferSize sets the size of the buffer when reading from in port
// The default buffersize is 1024
func BufferSize(buffersize int) InOption {
	return func(i *in) {
		i.buffersize = buffersize
	}
}

// SleepingTime sets the duration for sleeping between reads when polling on in port
// The default sleeping time is 0.1ms
func SleepingTime(d time.Duration) InOption {
	return func(i *in) {
		i.sleepingTime = d
	}
}

// In wraps an input portmidi.Stream to make it compatible with gomidi/mid#In
func In(i *portmidi.Stream, options ...InOption) mid.In {
	_in := &in{Stream: i}
	_in.buffersize = 1024

	// sleepingTime of 0.1ms should be fine to prevent busy waiting
	// and still fast enough for performances
	_in.sleepingTime = time.Nanosecond * 1000 * 100

	for _, opt := range options {
		opt(_in)
	}

	return _in
}

type in struct {
	*portmidi.Stream
	mx           sync.Mutex
	stopped      bool
	buffersize   int
	sleepingTime time.Duration
}

// StopListening cancels the listening
func (i *in) StopListening() {
	i.mx.Lock()
	i.stopped = true
	i.mx.Unlock()
}

func (i *in) read(cb func([]byte)) error {
	events, err := i.Read(i.buffersize)

	if err != nil {
		return err
	}

	for _, ev := range events {
		var b = make([]byte, 3)
		b[0] = byte(ev.Status)
		b[1] = byte(ev.Data1)
		b[2] = byte(ev.Data2)
		cb(b)
	}

	return nil
}

// SetListener makes the listener listen to the in port
func (i *in) SetListener(f func([]byte)) {
	for i.stopped == false {
		i.read(f)
		time.Sleep(i.sleepingTime)
	}
}
