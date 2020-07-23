package templatemethodd

import "fmt"

//IAbstractMethod 抽象方法接口
type IAbstractMethod interface {
	PrimitiveOperation1()
	PrimitiveOperation2()
}

//ITemplateMethod 模板方法接口
type ITemplateMethod interface {
	TemplateMethod()
}

//Template 模板
type Template struct {
	IAbstractMethod
}

//NewTemplate Template构造函数
func NewTemplate(abstractMethod IAbstractMethod) *Template {
	return &Template{
		IAbstractMethod: abstractMethod,
	}
}

//TemplateMethod 模板的方法
func (t *Template) TemplateMethod() {
	t.IAbstractMethod.PrimitiveOperation1()
	t.IAbstractMethod.PrimitiveOperation2()
}

//ConcreteClassA 具体类A
type ConcreteClassA struct {
	*Template
}

//NewConcreteClassA ConcreteClassA构造函数
func NewConcreteClassA() ITemplateMethod {
	classA := &ConcreteClassA{}
	template := NewTemplate(classA)
	classA.Template = template
	return classA
}

//PrimitiveOperation1 方法1
func (*ConcreteClassA) PrimitiveOperation1() {
	fmt.Println("具体类A方法1实现")
}

//PrimitiveOperation2 方法2
func (*ConcreteClassA) PrimitiveOperation2() {
	fmt.Println("具体类A方法2实现")
}

//ConcreteClassB 具体类B
type ConcreteClassB struct {
	*Template
}

//NewConcreteClassB ConcreteClassB构造函数
func NewConcreteClassB() ITemplateMethod {
	classB := &ConcreteClassB{}
	template := NewTemplate(classB)
	classB.Template = template
	return classB
}

//PrimitiveOperation1 方法1
func (*ConcreteClassB) PrimitiveOperation1() {
	fmt.Println("具体类B方法1实现")
}

//PrimitiveOperation2 方法2
func (*ConcreteClassB) PrimitiveOperation2() {
	fmt.Println("具体类B方法2实现")
}
