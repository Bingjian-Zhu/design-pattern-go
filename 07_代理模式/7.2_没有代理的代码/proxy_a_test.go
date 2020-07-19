package proxya

func ExampleProxy_a() {
	jiaojiao := &SchoolGirl{Name: "李娇娇"}
	zhuojiayi := &Pursuit{mm: jiaojiao}
	zhuojiayi.GiveDolls()
	zhuojiayi.GiveFlowers()
	zhuojiayi.GiveChocolate()
	// Output:
	// 李娇娇 送你洋娃娃
	// 李娇娇 送你鲜花
	// 李娇娇 送你巧克力
}
