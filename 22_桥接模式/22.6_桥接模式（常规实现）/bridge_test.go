package bridge

func ExampleBridge() {
	ab := &RefinedAbstraction{}

	ab.SetImplementor(&ConcreteImplementorA{})
	ab.Operation()

	ab.SetImplementor(&ConcreteImplementorB{})
	ab.Operation()

	// OutPut:
	// 具体实现A的方法执行
	// 具体实现B的方法执行
}
