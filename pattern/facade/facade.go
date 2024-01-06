package main

import "fmt"

type Subsystem1 struct{}

func NewSubsystem1() (subsystem1 *Subsystem1) {
	subsystem1 = &Subsystem1{}
	return
}

func (s *Subsystem1) operation1() {
	fmt.Println("Subsystem 1: operation 1")
}

func (s *Subsystem1) operation2() {
	fmt.Println("Subsystem 1: operation 2")
}

type Subsystem2 struct{}

func NewSubsystem2() (subsystem2 *Subsystem2) {
	subsystem2 = &Subsystem2{}
	return
}

func (s *Subsystem2) operation3() {
	fmt.Println("Subsystem 2: operation 3")
}

func (s *Subsystem2) operation4() {
	fmt.Println("Subsystem 2: operation 4")
}

type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() (facade *Facade) {
	facade = &Facade{}
	facade.subsystem1 = NewSubsystem1()
	facade.subsystem2 = NewSubsystem2()
	return
}

func (facade *Facade) facadeOperation1() {
	facade.subsystem1.operation1()
}

func (facade *Facade) facadeOperation2() {
	facade.subsystem1.operation2()
	facade.subsystem2.operation3()
}

func main() {
	facade := NewFacade()
	facade.facadeOperation1()
	facade.facadeOperation2()
}
