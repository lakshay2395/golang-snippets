package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training!!!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

func main() {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}
