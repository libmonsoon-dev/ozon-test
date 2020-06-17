package main

import "sync"

type Listener chan struct{}
type Listeners map[Listener]struct{}

type TreadSafeMap struct {
	data         map[int]int
	mu           sync.Mutex
	listenersMap map[int]Listeners
}

func NewTreadSafeMap(size int) *TreadSafeMap {
	return &TreadSafeMap{
		data:         make(map[int]int, size),
		listenersMap: make(map[int]Listeners),
	}
}

func (tsm *TreadSafeMap) Store(key, value int) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()

	tsm.data[key] = value

	if listenersSet, ok := tsm.listenersMap[key]; ok {
		for listener := range listenersSet {
			listener <- struct{}{}
		}
	}
}

func (tsm *TreadSafeMap) MustLoadAndDelete(key int) (value int) {
	var ok bool
	value, ok = tsm.LoadAndDelete(key)

	if ok {
		return
	}

	tsm.WaitForKey(key)
	value, ok = tsm.LoadAndDelete(key)
	if !ok {
		panic("Already deleted")
	}
	return
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

func (tsm *TreadSafeMap) WaitForKey(key int) {
	if listener := tsm.CreateListener(key); listener != nil {
		<-listener
		tsm.RemoveListener(key, listener)
	}
}

func (tsm *TreadSafeMap) CreateListener(key int) Listener {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()

	if _, ok := tsm.data[key]; ok {
		return nil
	}

	listener := make(Listener)
	if _, ok := tsm.listenersMap[key]; !ok {
		tsm.listenersMap[key] = make(Listeners, 1)
	}

	tsm.listenersMap[key][listener] = struct{}{}

	return listener
}

func (tsm *TreadSafeMap) RemoveListener(key int, listener Listener) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()

	delete(tsm.listenersMap[key], listener)

	if len(tsm.listenersMap[key]) == 0 {
		delete(tsm.listenersMap, key)
	}
}
