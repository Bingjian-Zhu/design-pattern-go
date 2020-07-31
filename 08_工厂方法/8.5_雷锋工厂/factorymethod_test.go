package factorymethod

func ExampleFactorymethod() {
	//简单工厂实现
	simpleFactory := &SimpleFactory{}
	studentA := simpleFactory.CreateLeiFeng("学雷锋的大学生")
	studentA.BuyRice()
	studentB := simpleFactory.CreateLeiFeng("学雷锋的大学生")
	studentB.Sweep()
	studentC := simpleFactory.CreateLeiFeng("学雷锋的大学生")
	studentC.Wash()
	// OutPut:
	// 买米
	// 扫地
	// 洗衣

	//工厂方法实现
	factory := &UndergraduateFactory{}
	student := factory.CreateLeiFeng()
	student.BuyRice()
	student.Sweep()
	student.Wash()
	// 买米
	// 扫地
	// 洗衣
}
