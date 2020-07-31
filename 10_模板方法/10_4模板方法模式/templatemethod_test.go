package templatemethod

func ExampleTemplatemethod() {
	c := NewConcreteClassA()
	c.TemplateMethod()

	c = NewConcreteClassB()
	c.TemplateMethod()

	// OutPut:
	// 具体类A方法1实现
	// 具体类A方法2实现
	// 具体类B方法1实现
	// 具体类B方法2实现
}
