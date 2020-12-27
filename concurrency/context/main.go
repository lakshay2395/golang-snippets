package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func generator(ctx context.Context) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("base err = ", ctx.Err())
				return
			case c <- rand.Intn(100):
			}
		}
	}()
	return c
}

func repeat(ctx context.Context, dataStream <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for c := range dataStream {
			select {
			case <-ctx.Done():
				fmt.Println("child err = ", ctx.Err())
				return
			case ch <- c:
			}
		}
	}()
	return ch
}

func main() {
	ctx := context.Background()
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	stream := generator(ctx2)
	result := repeat(ctx2, stream)
	for r := range result {
		fmt.Println("r = ", r)
		cancel()
	}
}
