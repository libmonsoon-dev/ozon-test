package main

func NewSemaphore(size int) Semaphore {
	if size < 1 {
		panic("Size must be 1 or grater")
	}
	return make(Semaphore, size)
}

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}
