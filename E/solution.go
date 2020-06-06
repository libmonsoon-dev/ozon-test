package main

import (
	"runtime"
	"sync"
)

const (
	maxGoroutinesPerWorker = 100
)

type TreadSafeCounter struct {
	value int
	mu    sync.RWMutex
}

func (c *TreadSafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *TreadSafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.value
}

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
	currentJobId := new(TreadSafeCounter)
	semaphore := make(chan struct{}, maxGoroutinesPerWorker)

	for i := 0; i < n; i++ {
		semaphore <- struct{}{}
		go run(fn, <-input, output, i, currentJobId, semaphore)
	}
}

func run(fn func(int) int, arg int, output chan<- int, jobId int, currentJobId *TreadSafeCounter, semaphore <-chan struct{}) {
	result := fn(arg)

	for currentJobId.Value() != jobId {
		runtime.Gosched()
	}

	output <- result
	currentJobId.Increment()
	<-semaphore
}
