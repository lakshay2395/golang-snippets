package main

import "sync"

func fanIn(done <-chan int, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedChannel := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for value := range c {
			select {
			case <-done:
				return
			case multiplexedChannel <- value:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexedChannel)
	}()
	return multiplexedChannel
}

func main() {

}
