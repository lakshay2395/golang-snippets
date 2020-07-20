package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	lock  sync.Mutex
	count int
}

func (m *MutexCounter) Increment() {
	m.lock.Lock()
	m.count++
	fmt.Println("counter after increment = ", m.count)
	m.lock.Unlock()
}

func (m *MutexCounter) Decrement() {
	m.lock.Lock()
	m.count--
	fmt.Println("counter after increment = ", m.count)
	m.lock.Unlock()
}

func (m *MutexCounter) Get() int {
	m.lock.Lock()
	value := m.count
	m.lock.Unlock()
	return value
}
