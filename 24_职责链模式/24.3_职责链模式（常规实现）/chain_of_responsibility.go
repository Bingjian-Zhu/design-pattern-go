package chain

import (
	"fmt"
	"reflect"
)

type IHandler interface {
	HandleRequest(request int)
}

type Handler struct {
	IHandler
	successor *Handler
}

func (h *Handler) HandleRequest(request int) {
	h.IHandler.HandleRequest(request)
	if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *Handler) SetSuccessor(s *Handler) {
	h.successor = s
}

type ConcreteHandler1 struct {
}

func NewConcreteHandler1() *Handler {
	return &Handler{
		IHandler: new(ConcreteHandler1),
	}
}

func (c *ConcreteHandler1) HandleRequest(request int) {
	if request >= 0 && request < 10 {
		fmt.Printf("%s  处理请求  %d\n", reflect.TypeOf(c).Elem().Name(), request)
	}
}

type ConcreteHandler2 struct {
}

func NewConcreteHandler2() *Handler {
	return &Handler{
		IHandler: new(ConcreteHandler2),
	}
}

func (c *ConcreteHandler2) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		fmt.Printf("%s  处理请求  %d\n", reflect.TypeOf(c).Elem().Name(), request)
	}
}

type ConcreteHandler3 struct {
}

func NewConcreteHandler3() *Handler {
	return &Handler{
		IHandler: new(ConcreteHandler3),
	}
}

func (c *ConcreteHandler3) HandleRequest(request int) {
	if request >= 20 && request < 30 {
		fmt.Printf("%s  处理请求  %d\n", reflect.TypeOf(c).Elem().Name(), request)
	}
}
