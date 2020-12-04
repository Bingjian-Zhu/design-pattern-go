package bridge

func ExampleBridge() {
	var ab HandsetSoft
	ab = new(HandsetBrandMGame)
	ab.Run()

	ab = new(HandsetBrandNGame)
	ab.Run()

	ab = new(HandsetBrandMAddressList)
	ab.Run()

	ab = new(HandsetBrandNAddressList)
	ab.Run()

	// 	OutPut:
	// 运行M品牌手机游戏
	// 运行N品牌手机游戏
	// 运行M品牌手机通讯录
	// 运行N品牌手机通讯录

}
