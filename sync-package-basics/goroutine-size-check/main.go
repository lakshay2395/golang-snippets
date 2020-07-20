package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

func main() {
	var c chan interface{}
	var wg sync.WaitGroup

	noob := func() { wg.Done(); <-c }
	numberOfGoroutines := 1e4
	wg.Add(int(numberOfGoroutines))
	before := memConsumed()
	for i := numberOfGoroutines; i > 0; i-- {
		go noob()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numberOfGoroutines/1000)
}
