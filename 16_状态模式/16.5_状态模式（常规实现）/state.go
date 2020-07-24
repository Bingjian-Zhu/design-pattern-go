package state

import (
	"fmt"
	"reflect"
)

//State 状态接口
type State interface {
	Handle(context *Context)
}

//ConcreteStateA 具体状态A
type ConcreteStateA struct {
}

//Handle 处理状态
func (*ConcreteStateA) Handle(context *Context) {
	context.SetState(&ConcreteStateB{})
}

//ConcreteStateB 具体状态B
type ConcreteStateB struct {
}

//Handle 处理状态
func (*ConcreteStateB) Handle(context *Context) {
	context.SetState(&ConcreteStateA{})
}

type Context struct {
	State
}

func NewContext(state State) *Context {
	return &Context{
		State: state,
	}
}

func (c *Context) GetState() State {
	return c.State
}

func (c *Context) SetState(state State) {
	c.State = state
	fmt.Printf("当前状态:%s\n", reflect.TypeOf(state).Elem().Name())
}

func (c *Context) Request() {
	c.State.Handle(c)
}
