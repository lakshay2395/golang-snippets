package main

import "fmt"

func orDone(done, channel <-chan int) <-chan int {
	stream := make(chan int)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
			case v, ok := <-channel:
				if !ok {
					return
				}
				select {
				case stream <- v:
				case <-done:
				}
			}
		}
	}()
	return stream
}

func bridge(done <-chan int, channel <-chan <-chan int) <-chan int {
	stream := make(chan int)
	go func() {
		for {
			var cha <-chan int
			select {
			case <-done:
				return
			case ch, ok := <-channel:
				if !ok {
					return
				}
				cha = ch
			}
			for data := range orDone(done, cha) {
				select {
				case stream <- data:
				case <-done:
					return
				}
			}
		}
	}()
	return stream
}

func main() {
	genVals := func() <-chan <-chan int {
		chanStream := make(chan (<-chan int))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan int, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}
	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
