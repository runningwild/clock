package clock

import (
	"github.com/runningwild/clock/core"
	"time"
)

// Clock models some of the basic functions from the standard time package.
type Clock interface {
	Now() time.Time
	At(t time.Time) <-chan time.Time
	After(d time.Duration) <-chan time.Time
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
}

type FakeClock struct {
	clock core.FakeTimeObj
}

func (f *FakeClock) Now() time.Time {
	return time.Unix(0, f.clock.GetCurrentTime())
}
func (f *FakeClock) At(t time.Time) <-chan time.Time {
	c := make(chan time.Time)
	go func() {
		c <- time.Unix(0, <-f.clock.At(t.UnixNano()))
	}()
	return c
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
func (f *FakeClock) Inc(d time.Duration) {
	f.clock.Inc(int64(d))
}

type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now()
}
func (RealClock) At(t time.Time) <-chan time.Time {
	return time.After(t.Sub(time.Now()))
}
func (RealClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
func (RealClock) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (RealClock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
