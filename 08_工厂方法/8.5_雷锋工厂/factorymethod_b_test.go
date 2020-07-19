package factorymethodb

func ExampleFactorymethod_b() {
	factory := &UndergraduateFactory{}
	student := factory.CreateLeiFeng()
	student.LeiFeng.BuyRice()
	student.LeiFeng.Sweep()
	student.LeiFeng.Wash()
	// OutPut:
	// 买米
	// 扫地
	// 洗衣
}
