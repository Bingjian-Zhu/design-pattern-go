package decorator

func ExampleDecorator() {
	c := &ConcreteComponent{}
	d1 := &ConcreteDecoratorA{}
	d2 := &ConcreteDecoratorB{}

	d1.Component = c
	d2.Component = d1

	d2.Operation()

	// Output:
	// 具体对象的操作
	// 具体装饰对象A的操作
	// 具体装饰对象B的操作

}
