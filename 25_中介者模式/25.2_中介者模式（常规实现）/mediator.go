package mediator

import "fmt"

//Mediator 中介者接口
type Mediator interface {
	Send(message string, colleague Colleague)
}

//ConcreteMediator 具体中介者
type ConcreteMediator struct {
	Colleague1 *ConcreteColleague1
	Colleague2 *ConcreteColleague2
}

//Send 中介者转发消息
func (c *ConcreteMediator) Send(message string, colleague Colleague) {
	if colleague == c.Colleague1 {
		c.Colleague2.Notify(message)
	} else {
		c.Colleague1.Notify(message)
	}
}

//Colleague 消息接收者
type Colleague interface {
	Send(message string)
	Notify(message string)
}

//ConcreteColleague1 同事1
type ConcreteColleague1 struct {
	mediator Mediator
}

//NewConcreteColleague1 同事1构造函数
func NewConcreteColleague1(mediator Mediator) *ConcreteColleague1 {
	return &ConcreteColleague1{
		mediator: mediator,
	}
}

//Send 发消息
func (c *ConcreteColleague1) Send(message string) {
	c.mediator.Send(message, c)
}

//Notify 通知
func (c *ConcreteColleague1) Notify(message string) {
	fmt.Println("同事1得到信息:" + message)
}

//ConcreteColleague2 同事2
type ConcreteColleague2 struct {
	mediator Mediator
}

//NewConcreteColleague2 同事2构造函数
func NewConcreteColleague2(mediator Mediator) *ConcreteColleague2 {
	return &ConcreteColleague2{
		mediator: mediator,
	}
}

//Send 发消息
func (c *ConcreteColleague2) Send(message string) {
	c.mediator.Send(message, c)
}

//Notify 通知
func (c *ConcreteColleague2) Notify(message string) {
	fmt.Println("同事2得到信息:" + message)
}
