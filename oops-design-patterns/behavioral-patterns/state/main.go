package main

import "fmt"

type PackageState interface {
	Next(*Package)
	Prev(*Package)
	PrintStatus()
}

type Package struct {
	state PackageState
}

func (p *Package) PreviousState() {
	p.state.Prev(p)
}

func (p *Package) NextState() {
	p.state.Next(p)
}

func (p *Package) PrintStatus() {
	p.state.PrintStatus()
}

type OrderedPackageState struct{}

type DeliveredPackageState struct{}

type ReceivedPackageState struct{}

func (o *OrderedPackageState) Next(pkg *Package) {
	pkg.state = &DeliveredPackageState{}
}

func (o *OrderedPackageState) Prev(pkg *Package) {
	fmt.Println("Intial stage")
}

func (o *OrderedPackageState) PrintStatus() {
	fmt.Println("ordered status")
}

func (o *DeliveredPackageState) Next(pkg *Package) {
	pkg.state = &ReceivedPackageState{}
}

func (o *DeliveredPackageState) Prev(pkg *Package) {
	pkg.state = &OrderedPackageState{}
}

func (o *DeliveredPackageState) PrintStatus() {
	fmt.Println("delivered status")
}

func (o *ReceivedPackageState) Next(pkg *Package) {
	fmt.Println("This package is already received by a client.")
}

func (o *ReceivedPackageState) Prev(pkg *Package) {
	pkg.state = &DeliveredPackageState{}
}

func (o *ReceivedPackageState) PrintStatus() {
	fmt.Println("received status")
}
