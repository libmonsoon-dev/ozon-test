package main

import (
	"sync"
)

const (
	maxGoroutinesPerWorker = 100
)

func Merge2Channels(fn func(int) int, in1, in2 <-chan int, out chan<- int, n int) {
	go merge2Channels(fn, in1, in2, out, n)
}

func merge2Channels(fn func(int) int, input1, input2 <-chan int, out chan<- int, n int) {
	result1, result2 := make(chan int), make(chan int)

	go newWorker(fn, input1, result1, n).Run()
	go newWorker(fn, input2, result2, n).Run()

	for i := 0; i < n; i++ {
		sum := <-result1 + <-result2
		out <- sum
	}
}

func newWorker(fn func(int) int, input <-chan int, output chan<- int, n int) *worker {
	semaphore := NewSemaphore(maxGoroutinesPerWorker)

	return &worker{
		fn:        fn,
		input:     input,
		output:    output,
		n:         n,
		semaphore: semaphore,

		cond: sync.Cond{L: new(sync.Mutex)},
	}
}

type worker struct {
	fn        func(int) int
	input     <-chan int
	output    chan<- int
	n         int
	semaphore Semaphore

	cond         sync.Cond
	currentJobId int
}

func (w *worker) Run() {
	for i := 0; i < w.n; i++ {
		w.semaphore.Acquire()
		go w.run(<-w.input, i)
	}
}

func (w *worker) run(arg int, jobId int) {
	defer w.semaphore.Release()

	result := w.fn(arg)

	w.cond.L.Lock()
	defer w.cond.L.Unlock()

	for jobId != w.currentJobId {
		w.cond.Wait()
	}

	w.output <- result
	w.currentJobId++
	w.cond.Broadcast()
}
