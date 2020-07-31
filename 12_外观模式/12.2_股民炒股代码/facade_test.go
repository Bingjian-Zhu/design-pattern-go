package facade

func ExampleFacade() {
	gu1 := &Stock1{}
	gu2 := &Stock2{}
	gu3 := &Stock3{}
	nd1 := &NationalDebt1{}
	rt1 := &Realty1{}

	gu1.Buy()
	gu2.Buy()
	gu3.Buy()
	nd1.Buy()
	rt1.Buy()

	gu1.Sell()
	gu2.Sell()
	gu3.Sell()
	nd1.Sell()
	rt1.Sell()

	// OutPut:
	// 股票1买入
	// 股票2买入
	// 股票3买入
	// 国债1买入
	// 房产1买入
	// 股票1卖出
	// 股票2卖出
	// 股票3卖出
	// 国债1卖出
	// 房产1卖出
}
