package decorator

import "fmt"

func ExampleDecorator() {
	xc := NewConcretePerson("小菜")
	fmt.Println("\n第一种装扮：")

	pqx := new(Sneakers)
	kk := new(BigTrouser)
	dtx := new(TShirts)

	pqx.Person = xc
	kk.Person = pqx
	dtx.Person = kk
	dtx.Show()

	fmt.Println("\n第二种装扮：")

	px := new(LeatherShoes)
	ld := new(Tie)
	xz := new(Suit)

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
