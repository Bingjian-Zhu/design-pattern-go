package state

func ExampleState() {
	c := NewContext(new(ConcreteStateA))
	c.Request()
	c.Request()
	c.Request()
	c.Request()

	// OutPut:
	// 当前状态:ConcreteStateB
	// 当前状态:ConcreteStateA
	// 当前状态:ConcreteStateB
	// 当前状态:ConcreteStateA
}
