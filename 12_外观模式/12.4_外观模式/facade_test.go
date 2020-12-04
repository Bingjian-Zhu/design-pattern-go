package facade

func ExampleFacade() {
	facade := new(Facade)

	facade.MethodA()
	facade.MethodB()

	// OutPut:
	// 方法组A()----
	// 子系统方法一
	// 子系统方法二
	// 子系统方法四
	// 方法组B()----
	// 子系统方法二
	// 子系统方法三
}
