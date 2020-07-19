package proxyc

func ExampleProxy_c() {
	jiaojiao := &SchoolGirl{Name: "李娇娇"}
	zhuojiayi := &Pursuit{mm: jiaojiao}
	daili := &Proxy{gg: zhuojiayi}
	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
	// Output:
	// 李娇娇 送你洋娃娃
	// 李娇娇 送你鲜花
	// 李娇娇 送你巧克力
}
