package visitor

import (
	"fmt"
	"reflect"
)

//Visitor 访问者接口
type Visitor interface {
	VisitConcreteElementA(concreteElementA *ConcreteElementA)
	VisitConcreteElementB(concreteElementB *ConcreteElementB)
}

//ConcreteVisitor1 具体访问者
type ConcreteVisitor1 struct {
}

//VisitConcreteElementA 访问节点
func (v *ConcreteVisitor1) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Printf("%s被%s访问\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(v).Elem().Name())
}

//VisitConcreteElementB 访问节点
func (v *ConcreteVisitor1) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Printf("%s被%s访问\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(v).Elem().Name())
}

//ConcreteVisitor2 具体访问者
type ConcreteVisitor2 struct {
}

//VisitConcreteElementA 访问节点
func (c *ConcreteVisitor2) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Printf("%s被%s访问\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(c).Elem().Name())
}

//VisitConcreteElementB 访问节点
func (c *ConcreteVisitor2) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Printf("%s被%s访问\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(c).Elem().Name())
}

//Element 节点接口
type Element interface {
	Accept(visitor Visitor)
}

//ConcreteElementA 具体节点
type ConcreteElementA struct {
}

//Accept 接收函数
func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

func (c *ConcreteElementA) OperationA() {}

type ConcreteElementB struct {
}

//Accept 接收函数
func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}

func (c *ConcreteElementB) OperationB() {

}

//ObjectStructure 对象结构
type ObjectStructure struct {
	elements []Element
}

//Attach 添加
func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

//Detach 移除
func (o *ObjectStructure) Detach(element Element) {
	for index, val := range o.elements {
		if val == element {
			o.elements = append(o.elements[:index], o.elements[:index]...)
			break
		}
	}
}

//Accept 接收
func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, e := range o.elements {
		e.Accept(visitor)
	}
}
