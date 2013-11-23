package clock_test

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
	"github.com/runningwild/clock/core"
	"sort"
)

func FakeTimeSpec(c gospec.Context) {
	c.Specify("time.After() will wake goroutines up in the right order.", func() {
		time := core.FakeTimeObj{}
		n := []int{500, 20, 34, 1012, 45, 3, 26, 76}
		s := make([]int, len(n))
		copy(s, n)
		sort.Ints(s)
		collect := make(chan int64)
		for _, v := range n {
			go func(v int64, after <-chan int64) {
				<-after
				collect <- v
			}(int64(v), time.After(int64(v)))
		}
		var prev int64
		for _, v := range s {
			time.Inc(int64(v) - prev)
			prev = int64(v)
			c.Expect(<-collect, Equals, int64(v))
		}
	})
}
