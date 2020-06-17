package main

import (
	"fmt"
	"testing"
	"time"
)


func fib(n int) int {
	switch n {
	case 0, 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func sleep(n int) int {
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func send(ch chan<- int, n, times int) {
	for i := 0; i < times; i++ {
		ch <- n
	}
}

func receive(ch <-chan int, times int) {
	for i := 0; i < times; i++ {
		<-ch
	}
}

func BenchmarkMerge2Channels(b *testing.B) {

	for _, bench := range []struct {
		arg int
		fn  func(int) int
	}{
		{1, fib},
		{3, fib},
		{5, fib},
		{15, fib},
		{20, fib},
		{25, fib},
		{30, fib},
		{1, sleep},
		{10, sleep},
		{100, sleep},
		{500, sleep},
	} {
		b.Run(fmt.Sprintf("Send \"%v\" b.N times", bench.arg), func(b *testing.B) {
			input1, input2, output := make(chan int), make(chan int), make(chan int)

			go send(input1, bench.arg, b.N)
			go send(input2, bench.arg, b.N)
			b.ReportAllocs()

			Merge2Channels(bench.fn, input1, input2, output, b.N)
			receive(output, b.N)

		})
	}

}
