package main

import "fmt"

type Command interface {
	Execute() string
}

type ConcreteCommand struct {
	receiver Receiver
}

func (cc *ConcreteCommand) Execute() string {
	return cc.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}

type Receiver struct{}

func (r *Receiver) Action() string {
	return "Действие выполнено"
}

func main() {
	receiver := &Receiver{}
	concreteCommand := &ConcreteCommand{receiver: *receiver}
	invoker := &Invoker{command: concreteCommand}
	result := invoker.ExecuteCommand()
	fmt.Println(result)
}
