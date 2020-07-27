package bridge

import "fmt"

type IAbstraction interface {
	Operation()
}

type Abstraction struct {
	implementor Implementor
}

func (a *Abstraction) SetImplementor(i Implementor) {
	a.implementor = i
}

func (a *Abstraction) Operation() {
	a.implementor.Operation()
}

type RefinedAbstraction struct {
	Abstraction
}

func (r *RefinedAbstraction) Operation() {
	r.Abstraction.Operation()
}

type Implementor interface {
	Operation()
}

type ConcreteImplementorA struct {
}

func (*ConcreteImplementorA) Operation() {
	fmt.Println("具体实现A的方法执行")
}

type ConcreteImplementorB struct {
}

func (*ConcreteImplementorB) Operation() {
	fmt.Println("具体实现B的方法执行")
}
