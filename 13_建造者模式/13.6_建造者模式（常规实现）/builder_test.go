package build

func ExampleBuild() {
	director := &Director{}

	b1 := NewConcreteBuilder1()
	director.Construct(b1)
	p1 := b1.GetResult()
	p1.Show()

	b2 := NewConcreteBuilder2()
	director.Construct(b2)
	p2 := b2.GetResult()
	p2.Show()

	// OutPut:
	// 产品 创建 ----
	// 部件A
	// 部件B
	// 产品 创建 ----
	// 部件X
	// 部件Y
}
