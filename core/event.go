package core

import (
	"sync"
)

type Event struct {
	time  int64
	mutex sync.Mutex
}

type EventHeap []*Event

func (h EventHeap) Push(e *Event) {
	h = append(h, e)
	h.up(len(h) - 1)
}

func (h EventHeap) Pop() *Event {
	n := len(h) - 1
	v := h[n]
	h[0], h[n] = h[n], h[0]
	h.down(0, n)
	h = h[0:n]
	return v
}

func (h EventHeap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(h[j].time < h[i].time) {
			break
		}
		h[i], h[j] = h[j], h[i]
		j = i
	}
}

func (h EventHeap) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !(h[j1].time < h[j2].time) {
			j = j2 // = 2*i + 2  // right child
		}
		if !(h[j].time < h[i].time) {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
}
