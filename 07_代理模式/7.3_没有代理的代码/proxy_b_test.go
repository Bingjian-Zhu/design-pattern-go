package proxyb

func ExampleProxy_b() {
	jiaojiao := &SchoolGirl{Name: "李娇娇"}
	daili := &Proxy{mm: jiaojiao}
	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
	// Output:
	// 李娇娇 送你洋娃娃
	// 李娇娇 送你鲜花
	// 李娇娇 送你巧克力
}
