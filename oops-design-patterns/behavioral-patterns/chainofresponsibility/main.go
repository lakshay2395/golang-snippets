package main

import (
	"fmt"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(msg string) {
	fmt.Printf("First logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)
		if se.NextChain != nil {
			se.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}
