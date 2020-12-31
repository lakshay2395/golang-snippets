package main

type Discounter interface {
	ApplyDiscount(price float32) float32
}

type EasterDiscounter struct{}

func (e *EasterDiscounter) ApplyDiscount(price float32) float32 {
	return price * 0.9
}

type ChristmasDiscounter struct{}

func (c *ChristmasDiscounter) ApplyDiscount(price float32) float32 {
	return price * 0.7
}

//use discounter interface and plug appropriate discounter for appropriate algo
