package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("func1 executed")
		wg.Done()
	}()
	go func() {
		fmt.Println("func2 executed")
		wg.Done()
	}()
	wg.Wait()
}
