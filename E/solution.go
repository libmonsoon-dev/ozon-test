package main

import (
	"runtime"
)

const (
	maxGoroutinesPerWorker = 100
)

func Merge2Channels(fn func(int) int, in1, in2 <-chan int, out chan<- int, n int) {
	go merge2Channels(fn, in1, in2, out, n)
}

func merge2Channels(fn func(int) int, input1, input2 <-chan int, out chan<- int, n int) {
	result1, result2 := make(chan int), make(chan int)

	go worker(fn, input1, result1, n)
	go worker(fn, input2, result2, n)

	for i := 0; i < n; i++ {
		sum := <-result1 + <-result2
		out <- sum
	}
}

func worker(fn func(int) int, input <-chan int, output chan<- int, n int) {
	semaphore := make(chan struct{}, maxGoroutinesPerWorker)

	data := NewTreadSafeMap(maxGoroutinesPerWorker)

	go func(data *TreadSafeMap) {
		for i := 0; i < n; i++ {
			for {
				if value, ok := data.LoadAndDelete(i); ok {
					output <- value
					break
				}

				runtime.Gosched()
			}
		}
	}(data)

	for i := 0; i < n; i++ {
		semaphore <- struct{}{}
		go run(fn, <-input, i, data, semaphore)
	}
}

func run(fn func(int) int, arg int, jobId int, data *TreadSafeMap, semaphore <-chan struct{}) {
	data.Store(jobId, fn(arg))
	<-semaphore
}
