package main

import "fmt"

func orDone(done,channel <-chan int) <-chan int{
	stream := make(chan int)
	go func(){
		defer close(stream)
		for {
			select{
			case <-done:
			case v,ok := <-channel:
				if !ok{
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

func main(){
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
	myChan := generator(done,1,2,3,4,5,6)
	for c := range orDone(done,myChan){
		fmt.Println(c)
	}
}