package main

import (
	"fmt"
	"io"
	"os"
)

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}
