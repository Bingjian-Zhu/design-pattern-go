package obserser

func ExampleObserser() {
	//前台小姐童子喆
	tongzizhe := &Secretary{}
	//看股票的同事
	tongshi1 := NewStockObserver("魏关姹", tongzizhe)
	//看NBA的同事
	tongshi2 := NewNBAObserver("易管查", tongzizhe)

	//前台记下了两位同事
	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)
	//发现老板回来
	tongzizhe.Action = "老板回来了！"
	//通知两个同事
	tongzizhe.Notify()

	// OutPut:
	// 老板回来了！ 魏关姹 关闭股票行情，继续工作！
	// 老板回来了！ 易管查 关闭NBA直播，继续工作！
}
