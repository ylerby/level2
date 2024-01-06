package main

import "fmt"

type IState interface {
	ActionOne()
	ActionTwo()
}

type State struct {
	context *Context
	IState
}

func NewState(iState IState, context *Context) (state *State) {
	state = &State{}
	state.IState = iState
	state.context = context

	context.SetState(state)
	return
}

type ConcreteStateOne struct {
	*State
}

func NewConcreteStateOne(context *Context) (concreteStateOne *ConcreteStateOne) {
	concreteStateOne = &ConcreteStateOne{}
	concreteStateOne.State = NewState(concreteStateOne, context)
	return
}

func (concreteStateOne *ConcreteStateOne) ActionOne() {
	fmt.Println("Вызов 'actionOne' - 'ConcreteStateOne'")
}

func (concreteStateOne *ConcreteStateOne) ActionTwo() {
	fmt.Println("Вызов 'actionTwo' - 'ConcreteStateOne'")
}

type ConcreteStateTwo struct {
	*State
}

func NewConcreteStateTwo(context *Context) (concreteStateTwo *ConcreteStateTwo) {
	concreteStateTwo = &ConcreteStateTwo{}
	concreteStateTwo.State = NewState(concreteStateTwo, context)
	return
}

func (concreteStateTwo *ConcreteStateTwo) ActionOne() {
	fmt.Println("Вызов 'actionOne' - 'ConcreteStateTwo'")
}

func (concreteStateTwo *ConcreteStateTwo) ActionTwo() {
	fmt.Println("Вызов 'actionTwo' - 'ConcreteStateTwo'")
}

type Context struct {
	state *State
}

func NewContext() (context *Context) {
	context = &Context{}
	return
}

func (context *Context) GetState() *State {
	return context.state
}

func (context *Context) PerformActionOne() {
	context.state.ActionOne()
}

func (context *Context) PerformActionTwo() {
	context.state.ActionTwo()
}

func (context *Context) SetState(state *State) {
	context.state = state
}

func main() {
	context := NewContext()
	NewConcreteStateOne(context)
	context.PerformActionOne()

	NewConcreteStateTwo(context)
	context.PerformActionOne()
	context.PerformActionTwo()
}
