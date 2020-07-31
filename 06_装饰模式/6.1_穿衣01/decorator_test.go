package decorator

import "fmt"

func ExampleDecorator() {
	xc := NewPerson("小菜")
	fmt.Println("\n第一种装扮：")
	xc.WearTShirts()
	xc.WearBigTrouser()
	xc.WearSneakers()
	xc.Show()

	fmt.Println("\n第二种装扮：")
	xc.WearSuit()
	xc.WearTie()
	xc.WearLeatherShoes()
	xc.Show()
	// Output:
	// 第一种装扮：
	// 大T恤 垮裤 破球鞋 装扮的小菜
	//
	// 第二种装扮：
	// 西装 领带 皮鞋 装扮的小菜
}
