package composite

import "fmt"

func ExampleComposite() {
	root := &ConcreteCompany{name: "北京总公司"}
	root.Add(&HRDepartment{name: "总公司人力资源部"})
	root.Add(&FinanceDepartment{name: "总公司财务部"})

	comp := &ConcreteCompany{name: "上海华东分公司"}
	comp.Add(&HRDepartment{name: "华东分公司人力资源部"})
	comp.Add(&FinanceDepartment{name: "华东分公司财务部"})
	root.Add(comp)

	comp1 := &ConcreteCompany{name: "南京办事处"}
	comp1.Add(&HRDepartment{name: "南京办事处人力资源部"})
	comp1.Add(&FinanceDepartment{name: "南京办事处财务部"})
	comp.Add(comp1)

	comp2 := &ConcreteCompany{name: "杭州办事处"}
	comp2.Add(&HRDepartment{name: "杭州办事处人力资源部"})
	comp2.Add(&FinanceDepartment{name: "杭州办事处财务部"})
	comp.Add(comp2)

	fmt.Println("结构图：")

	root.Display(1)

	fmt.Println("\n职责：")

	root.LineOfDuty()

	// 	OutPut:
	// 结构图：
	// -北京总公司
	// ---总公司人力资源部
	// ---总公司财务部
	// ---上海华东分公司
	// -----华东分公司人力资源部
	// -----华东分公司财务部
	// -----南京办事处
	// -------南京办事处人力资源部
	// -------南京办事处财务部
	// -----杭州办事处
	// -------杭州办事处人力资源部
	// -------杭州办事处财务部
	//
	// 职责：
	// 总公司人力资源部 员工招聘培训管理
	// 总公司财务部 公司财务收支管理
	// 华东分公司人力资源部 员工招聘培训管理
	// 华东分公司财务部 公司财务收支管理
	// 南京办事处人力资源部 员工招聘培训管理
	// 南京办事处财务部 公司财务收支管理
	// 杭州办事处人力资源部 员工招聘培训管理
	// 杭州办事处财务部 公司财务收支管理

}
