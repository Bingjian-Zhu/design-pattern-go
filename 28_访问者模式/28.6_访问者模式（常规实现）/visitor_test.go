package visitor

func ExampleVisitor() {
	o := &ObjectStructure{}
	o.Attach(&ConcreteElementA{})
	o.Attach(&ConcreteElementB{})

	v1 := &ConcreteVisitor1{}
	v2 := &ConcreteVisitor2{}

	o.Accept(v1)
	o.Accept(v2)

	// OutPut:
	// ConcreteElementA被ConcreteVisitor1访问
	// ConcreteElementB被ConcreteVisitor1访问
	// ConcreteElementA被ConcreteVisitor2访问
	// ConcreteElementB被ConcreteVisitor2访问
}
