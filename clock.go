package clock

import (
	"github.com/runningwild/clock/core"
	"time"
)

// Clock models some of the basic functions from the standard time package.
type Clock interface {
	After(d time.Duration) <-chan time.Time
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
}

type FakeClock struct {
	clock core.FakeTimeObj
}

func (f *FakeClock) After(d time.Duration) <-chan time.Time {
	c := make(chan time.Time)
	go func() {
		c <- time.Unix(0, <-f.clock.After(int64(d)))
	}()
	return c
}
func (f *FakeClock) Sleep(d time.Duration) {
	f.clock.Sleep(int64(d))
}
func (f *FakeClock) Tick(d time.Duration) <-chan time.Time {
	c := make(chan time.Time)
	go func() {
		ticker := f.clock.Tick(int64(d))
		for t := range ticker {
			c <- time.Unix(0, t)
		}
	}()
	return c
}

type RealClock struct{}

func (RealClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
func (RealClock) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (RealClock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
