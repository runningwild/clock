package core

import (
	"sync"
)

type FakeTimeObj struct {
	CurrentTime int64

	EventsMutex sync.Mutex
	Events      []Event
}

func (f *FakeTimeObj) pushEvent(t int64) *Event {
	f.Events = append(f.Events, Event{})
	return &f.Events[len(f.Events)-1]
}

func (f *FakeTimeObj) After(d int64) <-chan int64 {
	f.EventsMutex.Lock()
	defer f.EventsMutex.Unlock()
	Event := f.pushEvent(f.CurrentTime + d)
	Event.mutex.Lock()
	c := make(chan int64)
	go func() {
		Event.mutex.Lock()
		c <- f.CurrentTime
	}()
	return c
}

// func (f *FakeTimeObj) Sleep(d Duration) {
// 	<-f.After(d)
// }
