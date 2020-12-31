package main

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (m *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s \n", s)
	fmt.Println(newMsg)
	return newMsg
}

type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	var newMsg string
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return newMsg
}
