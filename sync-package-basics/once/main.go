package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var count int
	var once sync.Once
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}
	once.Do(increment)
	once.Do(decrement)
	fmt.Println("value of count = ", strconv.Itoa(count))

	deadlock()
}
