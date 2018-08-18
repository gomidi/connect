package portmidiadapter

import (
	"github.com/gomidi/mid"
	"github.com/rakyll/portmidi"
	"sync"
	"time"
)

func Out(o *portmidi.Stream) mid.Out {
	return &out{o}
}

type out struct {
	*portmidi.Stream
}

func (o *out) Send(b []byte) error {
	return o.WriteShort(int64(b[0]), int64(b[1]), int64(b[2]))
}

func In(i *portmidi.Stream) mid.In {
	return &in{Stream: i}
}

type in struct {
	*portmidi.Stream
	mx      sync.Mutex
	stopped bool
}

func (i *in) StopListening() {
	i.mx.Lock()
	i.stopped = true
	i.mx.Unlock()
}

func (i *in) read(cb func([]byte)) error {
	//1024
	//events, err := r.in.Read(3)
	events, err := i.Read(1024)

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

func (i *in) SetListener(f func([]byte)) {
	for i.stopped == false {
		i.read(f)
		// sleep 0.1ms, that should be fine to prevent busy waiting
		// and still fast enough for performances
		time.Sleep(time.Nanosecond * 1000 * 100)
	}
}
