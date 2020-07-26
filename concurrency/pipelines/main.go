package main

import "fmt"

func main() {
	generator := func(done <-chan int, integers ...int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range integers {
				select {
				case <-done:
					return
				case c <- i:
				}
			}
		}()
		return c
	}

	add := func(done <-chan int, additive int, stream <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range stream {
				select {
				case <-done:
					return
				case c <- i + additive:
				}
			}
		}()
		return c
	}

	multiply := func(done <-chan int, multiplier int, stream <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range stream {
				select {
				case <-done:
					return
				case c <- i * multiplier:
				}
			}
		}()
		return c
	}
	done := make(chan int)
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5, 6)
	result := multiply(done, 20, add(done, 10, intStream))
	for r := range result {
		fmt.Println(r)
	}
}
