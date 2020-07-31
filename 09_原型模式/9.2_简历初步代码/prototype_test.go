package prototype

func ExamplePrototype() {
	a := &Resume{name: "大鸟"}
	a.SetPersonalInfo("男", "29")
	a.SetWorkExperience("1998-2000", "XX公司")

	b := a
	c := a

	a.Display()
	b.Display()
	c.Display()
	// OutPut:
	// 大鸟 男 29
	// 工作经历：1998-2000 XX公司
	// 大鸟 男 29
	// 工作经历：1998-2000 XX公司
	// 大鸟 男 29
	// 工作经历：1998-2000 XX公司
}
