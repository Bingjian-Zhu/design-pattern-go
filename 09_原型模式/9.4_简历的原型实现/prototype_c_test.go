package prototypec

import (
	"fmt"
)

var manager *ConcretePrototype

//Resume 简历
type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

//SetPersonalInfo 设置个人信息
func (r *Resume) SetPersonalInfo(sex, age string) {
	r.sex = sex
	r.age = age
}

//SetWorkExperience 设置工作经历
func (r *Resume) SetWorkExperience(timeArea, company string) {
	r.timeArea = timeArea
	r.company = company
}

//Display 显示
func (r *Resume) Display() {
	fmt.Printf("%s %s %s\n", r.name, r.sex, r.age)
	fmt.Printf("工作经历：%s %s\n", r.timeArea, r.company)
}

func (r *Resume) Clone() Prototype {
	tc := *r
	return &tc
}

func ExamplePrototype_c() {
	manager = NewConcretePrototype()

	a := &Resume{
		name: "大鸟",
	}
	manager.Set("a", a)
	a.SetPersonalInfo("男", "29")
	a.SetWorkExperience("1998-2000", "XX公司")

	b := manager.Get("a").Clone().(*Resume)
	b.SetWorkExperience("1998-2006", "YY企业")

	c := manager.Get("a").Clone().(*Resume)
	a.Display()
	b.Display()
	c.Display()
	// OutPut:
	// 大鸟 男 29
	// 工作经历：1998-2000 XX公司
	// 大鸟 男 29
	// 工作经历：1998-2006 YY企业
	// 大鸟 男 29
	// 工作经历：1998-2000 XX公司
}
