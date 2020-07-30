package visitor

import (
	"fmt"
	"reflect"
)

//Action 接口
type Action interface {
	//得到男人结论或反应
	GetManConclusion(concreteElementA *Man)
	//得到女人结论或反应
	GetWomanConclusion(concreteElementB *Woman)
}

//Success 成功
type Success struct {
}

//GetManConclusion 得到男人结论或反应
func (s *Success) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%s%s时，背后多半有一个伟大的女人。\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(s).Elem().Name())
}

//GetWomanConclusion 得到女人结论或反应
func (s *Success) GetWomanConclusion(concreteElementB *Woman) {
	fmt.Printf("%s%s时，背后大多有一个不成功的男人。\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(s).Elem().Name())
}

//Failing 失败
type Failing struct {
}

//GetManConclusion 得到男人结论或反应
func (f *Failing) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%s%s时，闷头喝酒，谁也不用劝。\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(f).Elem().Name())
}

//GetWomanConclusion 得到女人结论或反应
func (f *Failing) GetWomanConclusion(concreteElementB *Woman) {
	fmt.Printf("%s%s时，眼泪汪汪，谁也劝不了。\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(f).Elem().Name())
}

//Amativeness 恋爱
type Amativeness struct {
}

//GetManConclusion 得到男人结论或反应
func (a *Amativeness) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%s%s时，凡事不懂也要装懂。\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(a).Elem().Name())
}

//GetWomanConclusion 得到女人结论或反应
func (a *Amativeness) GetWomanConclusion(concreteElementB *Woman) {
	fmt.Printf("%s%s时，遇事懂也装作不懂\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(a).Elem().Name())
}

//Marriage 结婚
type Marriage struct {
}

//GetManConclusion 得到男人结论或反应
func (m *Marriage) GetManConclusion(concreteElementA *Man) {
	fmt.Printf("%s%s时，感慨道：恋爱游戏终结时，‘有妻徒刑’遥无期。\n", reflect.TypeOf(concreteElementA).Elem().Name(), reflect.TypeOf(m).Elem().Name())
}

//GetWomanConclusion 得到女人结论或反应
func (m *Marriage) GetWomanConclusion(concreteElementB *Woman) {
	fmt.Printf("%s%s时，欣慰曰：爱情长跑路漫漫，婚姻保险保平安。\n", reflect.TypeOf(concreteElementB).Elem().Name(), reflect.TypeOf(m).Elem().Name())
}

//Person 人
type Person interface {
	//接受
	Accept(visitor Action)
}

//Man 男人
type Man struct {
}

//Accept 接受
func (m *Man) Accept(visitor Action) {
	visitor.GetManConclusion(m)
}

//Woman 女人
type Woman struct {
}

//Accept 接受
func (w *Woman) Accept(visitor Action) {
	visitor.GetWomanConclusion(w)
}

//ObjectStructure 对象结构
type ObjectStructure struct {
	elements []Person
}

//Attach 增加
func (o *ObjectStructure) Attach(element Person) {
	o.elements = append(o.elements, element)
}

//Detach 移除
func (o *ObjectStructure) Detach(element Person) {
	for index, val := range o.elements {
		if val == element {
			o.elements = append(o.elements[:index], o.elements[index+1:]...)
			break
		}
	}
}

//Display 查看显示
func (o *ObjectStructure) Display(visitor Action) {
	for _, e := range o.elements {
		e.Accept(visitor)
	}
}
