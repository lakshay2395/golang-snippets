package main

import (
	"fmt"
	"sync"
)

type RWMutexCounter struct {
	lock  sync.RWMutex
	count int
}

func (m *RWMutexCounter) Increment() {
	m.lock.Lock()
	m.count++
	fmt.Println("counter after increment = ", m.count)
	m.lock.Unlock()
}

func (m *RWMutexCounter) Decrement() {
	m.lock.Lock()
	m.count--
	fmt.Println("counter after increment = ", m.count)
	m.lock.Unlock()
}

//it provides a read lock; if none of the write operations are going on, we can read using any number of read locks
func (m *RWMutexCounter) Get() int {
	m.lock.RLock()
	value := m.count
	m.lock.RUnlock()
	return value
}
