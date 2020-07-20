package main

import "sync"

func deadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() {
		onceB.Do(initB)
	}
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
}
