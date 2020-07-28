package command

import "fmt"

type ICommand interface {
	Execute()
}

type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type Receiver struct {
}

func (*Receiver) Action() {
	fmt.Println("执行请求！")
}

type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
