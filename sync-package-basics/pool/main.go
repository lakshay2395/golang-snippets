package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance...")
			return struct{}{}
		},
	}
	myPool.Get() //new instance
	instance := myPool.Get() //new instance
	myPool.Put(instance) //instance returned
	myPool.Get() //old instance
	fmt.Println(runtime.NumCPU())
} 