package decoratorc

import (
	"fmt"
)

//Component 部件
type Component interface {
	Operation()
}

//ConcreteComponent 具体部件
type ConcreteComponent struct {
}

//Operation 具体操作
func (*ConcreteComponent) Operation() {
	fmt.Println("具体对象的操作")
}

//Decorator 装饰器
type Decorator struct {
	Component
}

//Operation 装饰器操作
func (decorator *Decorator) Operation() {
	if decorator.Component != nil {
		decorator.Component.Operation()
	}
}

//ConcreteDecoratorA 具体装饰对象A
type ConcreteDecoratorA struct {
	addedState string
	//想要调用父类方法，需要把父类组合到结构体中
	Component
}

//Operation 具体装饰对象A的操作
func (a *ConcreteDecoratorA) Operation() {
	a.Component.Operation()
	a.addedState = "New State"
	fmt.Println("具体装饰对象A的操作")
}

//ConcreteDecoratorB 具体装饰对象B
type ConcreteDecoratorB struct {
	//想要调用父类方法，需要把父类组合到结构体中
	Component
}

//Operation 具体装饰对象B的操作
func (b *ConcreteDecoratorB) Operation() {
	b.Component.Operation()
	b.AddedBehavior()
	fmt.Println("具体装饰对象B的操作")
}

//AddedBehavior 添加新特性
func (b *ConcreteDecoratorB) AddedBehavior() {

}
