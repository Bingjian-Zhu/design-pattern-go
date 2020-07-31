package prototype

import "fmt"

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
