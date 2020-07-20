package prototypea

func ExamplePrototype_a() {
	a := &Resume{name: "大鸟"}
	a.SetPersonalInfo("男", "29")
	a.SetWorkExperience("1998-2000", "XX公司")

	b := &Resume{name: "大鸟"}
	b.SetPersonalInfo("男", "29")
	b.SetWorkExperience("1998-2000", "XX公司")

	c := &Resume{name: "大鸟"}
	c.SetPersonalInfo("男", "29")
	c.SetWorkExperience("1998-2000", "XX公司")

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
