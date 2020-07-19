package decoratord

import "fmt"

func ExampleDecorator_a() {
	xc := NewConcretePerson("小菜")
	fmt.Println("\n第一种装扮：")

	pqx := &Sneakers{}
	kk := &BigTrouser{}
	dtx := &TShirts{}

	pqx.Person = xc
	kk.Person = pqx
	dtx.Person = kk
	dtx.Show()

	fmt.Println("\n第二种装扮：")

	px := &LeatherShoes{}
	ld := &Tie{}
	xz := &Suit{}

	px.Person = xc
	ld.Person = px
	xz.Person = ld
	xz.Show()

	// Output:
	// 第一种装扮：
	// 大T恤 垮裤 破球鞋 装扮的小菜
	//
	// 第二种装扮：
	// 西装 领带 皮鞋 装扮的小菜
}
