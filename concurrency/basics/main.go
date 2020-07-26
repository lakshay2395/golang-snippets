package main

import (
	"fmt"
	"sync"
)

//lexical scoping of go routines
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	chanOwner := func() <-chan int {
		results := make(chan int)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(ch <-chan int) {
		for c := range ch {
			fmt.Println(c)
		}
		wg.Done()
	}
	ch := chanOwner()
	consumer(ch)
	wg.Wait()
}
