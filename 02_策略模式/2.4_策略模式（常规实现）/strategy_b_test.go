package strategyb

func ExampleStrategy_b() {
	var context *Context
	context = NewContext(&ConcreteStrategyA{})
	context.ContextInterface()
	context = NewContext(&ConcreteStrategyB{})
	context.ContextInterface()
	context = NewContext(&ConcreteStrategyC{})
	context.ContextInterface()
	// Output:
	// 算法A实现
	// 算法B实现
	// 算法C实现
}
