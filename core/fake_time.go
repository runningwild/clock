package core

import (
	"sync"
	"sync/atomic"
)

type FakeTimeObj struct {
	CurrentTime int64

	EventsMutex sync.Mutex
	Events      EventHeap
}

func (f *FakeTimeObj) Inc(d int64) {
	f.EventsMutex.Lock()
	defer f.EventsMutex.Unlock()
	f.inc(d)
}

func (f *FakeTimeObj) GetCurrentTime() int64 {
	return atomic.LoadInt64(&f.CurrentTime)
}

func (f *FakeTimeObj) inc(d int64) {
	currentTime := atomic.AddInt64(&f.CurrentTime, d)
	for f.Events.Len() > 0 && f.Events[0].time <= currentTime {
		event := f.Events.Pop()
		event.mutex.Unlock()
	}
}

func (f *FakeTimeObj) At(t int64) <-chan int64 {
	f.EventsMutex.Lock()
	defer f.EventsMutex.Unlock()
	event := &Event{time: t}
	f.Events.Push(event)
	event.mutex.Lock()
	c := make(chan int64)
	go func() {
		event.mutex.Lock()
		c <- event.time
	}()
	f.inc(0)
	return c
}

func (f *FakeTimeObj) After(d int64) <-chan int64 {
	return f.At(f.GetCurrentTime() + d)
}

func (f *FakeTimeObj) Sleep(d int64) {
	<-f.After(d)
}

func (f *FakeTimeObj) Tick(d int64) <-chan int64 {
	c := make(chan int64)
	next := f.GetCurrentTime() + d
	go func() {
		for {
			<-f.At(next)
			c <- next
			next += d
		}
	}()
	return c
}
