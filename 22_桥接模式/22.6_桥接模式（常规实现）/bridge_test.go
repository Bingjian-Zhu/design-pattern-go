package bridge

func ExampleBridge() {
	ab := new(RefinedAbstraction)

	ab.SetImplementor(new(ConcreteImplementorA))
	ab.Operation()

	ab.SetImplementor(new(ConcreteImplementorB))
	ab.Operation()

	// OutPut:
	// 具体实现A的方法执行
	// 具体实现B的方法执行
}
