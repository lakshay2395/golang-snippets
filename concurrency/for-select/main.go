package main

import (
	"fmt"
)

func main() {
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
	ch := chanOwner()
	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println("Got data = ", data)
			} else {
				fmt.Println("channel closed")
				return
			}
		}
	}
}
