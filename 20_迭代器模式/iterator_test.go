package iterator

import "fmt"

func ExampleIterator() {
	a := &ConcreteAggregate{
		items: make([]interface{}, 6),
	}

	a.items[0] = "大鸟"
	a.items[1] = "小菜"
	a.items[2] = "行李"
	a.items[3] = "老外"
	a.items[4] = "公交内部员工"
	a.items[5] = "小偷"

	i := NewConcreteIterator(a)

	for !i.IsDone() {
		fmt.Print(i.CurrentItem())
		fmt.Println(" 请买车票!")
		i.Next()
	}

	// 	OutPut:
	// 大鸟 请买车票!
	// 小菜 请买车票!
	// 行李 请买车票!
	// 老外 请买车票!
	// 公交内部员工 请买车票!
	// 小偷 请买车票!

}
