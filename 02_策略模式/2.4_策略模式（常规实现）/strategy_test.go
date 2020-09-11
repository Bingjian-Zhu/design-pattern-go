package strategy

func ExampleStrategy() {
	var context *Context
	context = NewContext(new(ConcreteStrategyA))
	context.ContextInterface()
	context = NewContext(new(ConcreteStrategyB))
	context.ContextInterface()
	context = NewContext(new(ConcreteStrategyC))
	context.ContextInterface()
	// Output:
	// 算法A实现
	// 算法B实现
	// 算法C实现
}
