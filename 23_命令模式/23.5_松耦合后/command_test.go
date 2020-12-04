package command

func ExampleCommand() {
	//开店前的准备
	boy := new(Barbecuer)
	bakeMuttonCommand1 := NewBakeMuttonCommand(boy)
	bakeMuttonCommand2 := NewBakeMuttonCommand(boy)
	bakeChickenWingCommand1 := NewBakeChickenWingCommand(boy)
	girl := new(Waiter)

	//开门营业
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand2)
	girl.SetOrder(bakeChickenWingCommand1)

	girl.Notify()

	// OutPut:
	// 烤羊肉串!
	// 烤羊肉串!
	// 烤鸡翅!
}
