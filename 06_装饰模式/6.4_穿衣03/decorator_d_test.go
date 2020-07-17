package decoratord

import "fmt"

func ExampleDecorator_a() {
	xc := NewPerson("小菜")
	fmt.Println("\n第一种装扮：")

	pqx := &Sneakers{}
	kk := &BigTrouser{}
	dtx := &TShirts{}

	pqx.Decorate(xc)
	kk.Decorate(pqx)
	dtx.Decorate(kk)
	dtx.Show()

	fmt.Println("\n第二种装扮：")

	px := &LeatherShoes{}
	ld := &Tie{}
	xz := &Suit{}

	px.Decorate(xc)
	ld.Decorate(px)
	xz.Decorate(ld)
	xz.Show()

	fmt.Println("\n第三种装扮：")

	pqx2 := &Sneakers{}
	px2 := &LeatherShoes{}
	kk2 := &BigTrouser{}
	ld2 := &Tie{}

	pqx2.Decorate(xc)
	px2.Decorate(pqx)
	kk2.Decorate(px2)
	ld2.Decorate(kk2)

	ld2.Show()
	// Output:
	// 第一种装扮：
	// 大T恤 垮裤 破球鞋 装扮的小菜
	//
	// 第二种装扮：
	// 西装 领带 皮鞋 装扮的小菜
}
