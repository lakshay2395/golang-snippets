package main

import (
	"fmt"
	"sync"
)

//using waitgroup
func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello World Fixed")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
