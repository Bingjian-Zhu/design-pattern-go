package state

func ExampleState() {
	c := NewContext(&ConcreteStateA{})
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
