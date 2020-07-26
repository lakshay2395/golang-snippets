package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan int) <-chan int {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}
	orDone := make(chan int)
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			time.Sleep(after)
			fmt.Println("gave signal")
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(2*time.Minute),
		sig(2*time.Second),
	)
	fmt.Println(time.Since(start))
}
