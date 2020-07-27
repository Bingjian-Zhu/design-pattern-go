package composite

import (
	"fmt"
	"strings"
)

//Company 公司接口
type Company interface {
	Add(c Company)     //增加
	Remove(c Company)  //移除
	Display(depth int) //显示
	LineOfDuty()       //履行职责
}

//ConcreteCompany 具体公司
type ConcreteCompany struct {
	children []Company
	name     string
}

//Add 添加公司
func (con *ConcreteCompany) Add(c Company) {
	con.children = append(con.children, c)
}

//Remove 移除公司
func (con *ConcreteCompany) Remove(c Company) {
	for index, val := range con.children {
		if val == c {
			con.children = append(con.children[0:index], con.children[index+1:]...)
			break
		}
	}
}

//Display 展示
func (con *ConcreteCompany) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + con.name)
	for _, component := range con.children {
		component.Display(depth + 2)
	}
}

//LineOfDuty 履行职责
func (con *ConcreteCompany) LineOfDuty() {
	for _, component := range con.children {
		component.LineOfDuty()
	}
}

//HRDepartment 人力资源部
type HRDepartment struct {
	name string
}

//Add 添加公司
func (*HRDepartment) Add(c Company) {
}

//Remove 移除公司
func (*HRDepartment) Remove(c Company) {
}

//Display 展示
func (h *HRDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + h.name)
}

//LineOfDuty 履行职责
func (h *HRDepartment) LineOfDuty() {
	fmt.Printf("%s 员工招聘培训管理\n", h.name)
}

//FinanceDepartment 财务部
type FinanceDepartment struct {
	name string
}

//Add 添加公司
func (*FinanceDepartment) Add(c Company) {
}

//Remove 移除公司
func (*FinanceDepartment) Remove(c Company) {
}

//Display 展示
func (f *FinanceDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + f.name)
}

//LineOfDuty 履行职责
func (f *FinanceDepartment) LineOfDuty() {
	fmt.Printf("%s 公司财务收支管理\n", f.name)
}
