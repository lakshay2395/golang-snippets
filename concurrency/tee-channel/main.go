package main

import (
	"fmt"
	"strconv"
)

func tee(done, inputChannel <-chan int) (_, _ <-chan int) {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		defer close(c1)
		defer close(c2)
		for data := range inputChannel {
			var out1, out2 = c1, c2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case out1 <- data:
					out1 = nil
					fmt.Println("pushed", strconv.Itoa(data), "to c1")
				case out2 <- data:
					out2 = nil
					fmt.Println("pushed", strconv.Itoa(data), "to c2")
				}
			}
		}
	}()
	return c1, c2
}

func main() {
	generator := func(done <-chan int, integers ...int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for _, i := range integers {
				select {
				case <-done:
					return
				case c <- i:
				}
			}
		}()
		return c
	}
	done := make(chan int)
	stream := generator(done, 1, 2, 3, 4, 5, 6)
	c1, c2 := tee(done, stream)
	for {
		data, ok := <-c1
		if !ok {
			break
		} else {
			fmt.Println("Data = ", data, <-c2)
		}
	}
}
