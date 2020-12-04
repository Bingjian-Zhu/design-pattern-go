package visitor

func ExampleVisitor() {
	o := new(ObjectStructure)
	o.Attach(new(ConcreteElementA))
	o.Attach(new(ConcreteElementB))

	v1 := new(ConcreteVisitor1)
	v2 := new(ConcreteVisitor2)

	o.Accept(v1)
	o.Accept(v2)

	// OutPut:
	// ConcreteElementA被ConcreteVisitor1访问
	// ConcreteElementB被ConcreteVisitor1访问
	// ConcreteElementA被ConcreteVisitor2访问
	// ConcreteElementB被ConcreteVisitor2访问
}
