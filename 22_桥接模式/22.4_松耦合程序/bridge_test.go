package bridge

func ExampleBridge() {
	var ab HandsetSoft
	ab = &HandsetBrandMGame{}
	ab.Run()

	ab = &HandsetBrandNGame{}
	ab.Run()

	ab = &HandsetBrandMAddressList{}
	ab.Run()

	ab = &HandsetBrandNAddressList{}
	ab.Run()

	// 	OutPut:
	// 运行M品牌手机游戏
	// 运行N品牌手机游戏
	// 运行M品牌手机通讯录
	// 运行N品牌手机通讯录

}
