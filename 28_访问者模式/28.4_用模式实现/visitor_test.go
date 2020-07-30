package visitor

func ExampleVisitor() {
	o := &ObjectStructure{}
	o.Attach(&Man{})
	o.Attach(&Woman{})

	v1 := &Success{}
	o.Display(v1)

	v2 := &Failing{}
	o.Display(v2)

	v3 := &Amativeness{}
	o.Display(v3)

	v4 := &Marriage{}
	o.Display(v4)

	// OutPut:
	// ManSuccess时，背后多半有一个伟大的女人。
	// WomanSuccess时，背后大多有一个不成功的男人。
	// ManFailing时，闷头喝酒，谁也不用劝。
	// WomanFailing时，眼泪汪汪，谁也劝不了。
	// ManAmativeness时，凡事不懂也要装懂。
	// WomanAmativeness时，遇事懂也装作不懂
	// ManMarriage时，感慨道：恋爱游戏终结时，‘有妻徒刑’遥无期。
	// WomanMarriage时，欣慰曰：爱情长跑路漫漫，婚姻保险保平安。
}
