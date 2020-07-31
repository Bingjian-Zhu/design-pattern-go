package decorator

import "fmt"

func ExampleDecorator() {
	xc := NewPerson("小菜")
	fmt.Println("\n第一种装扮：")

	dtx := &TShirts{}
	kk := &BigTrouser{}
	pqx := &Sneakers{}

	dtx.Show()
	kk.Show()
	pqx.Show()
	xc.Show()

	fmt.Println("\n第二种装扮：")

	xz := &Suit{}
	ld := &Tie{}
	px := &LeatherShoes{}

	xz.Show()
	ld.Show()
	px.Show()
	xc.Show()
	// Output:
	// 第一种装扮：
	// 大T恤 垮裤 破球鞋 装扮的小菜
	//
	// 第二种装扮：
	// 西装 领带 皮鞋 装扮的小菜
}
