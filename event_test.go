package clock_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/runningwild/clock/core"
	"sort"
	"testing"
)

func FakeTimeTest(t testing.T) {
	Convey("time.After() wil wake go routines up in the right order", t, func() {
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
			So(<-collect, ShouldEqual, int64(v))
		}
	})
}
