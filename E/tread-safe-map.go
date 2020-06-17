package main

import "sync"

type TreadSafeMap struct {
	data map[int]int
	mu   sync.Mutex
}

func NewTreadSafeMap(size int) *TreadSafeMap {
	return &TreadSafeMap{data: make(map[int]int, size)}
}

func (tsm *TreadSafeMap) Store(key, value int) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()

	tsm.data[key] = value
}

func (tsm *TreadSafeMap) LoadAndDelete(key int) (value int, ok bool) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()

	value, ok = tsm.data[key]
	if ok {
		delete(tsm.data, key)
	}

	return
}
