package facade

func ExampleFacade() {
	jijin := new(Fund)

	jijin.BuyFund()
	jijin.SellFund()

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
