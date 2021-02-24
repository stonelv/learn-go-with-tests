package main

import "sync"

//InMemoryPlayerStore is the implementation of the PlayerStore
type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

//NewInMemoryPlayerStore create the object of the class
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

//GetPlayerScore is the implementation of the method in PlayerStore
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

//RecordWin used to save win info
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}
