package facadec

import "fmt"

//SubSystemOne 子系统一
type SubSystemOne struct {
}

//MethodOne 子系统方法一
func (*SubSystemOne) MethodOne() {
	fmt.Println("子系统方法一")
}

//SubSystemTwo 子系统二
type SubSystemTwo struct {
}

//MethodTwo 子系统方法二
func (*SubSystemTwo) MethodTwo() {
	fmt.Println("子系统方法二")
}

//SubSystemThree 子系统三
type SubSystemThree struct {
}

//MethodThree 子系统方法三
func (*SubSystemThree) MethodThree() {
	fmt.Println("子系统方法三")
}

//SubSystemFour 子系统四
type SubSystemFour struct {
}

//MethodFour 子系统方法四
func (*SubSystemFour) MethodFour() {
	fmt.Println("子系统方法四")
}

//Facade 外观
type Facade struct {
	one   *SubSystemOne
	two   *SubSystemTwo
	three *SubSystemThree
	four  *SubSystemFour
}

//NewFacade 构造函数
func NewFacade() *Facade {
	return &Facade{
		one:   &SubSystemOne{},
		two:   &SubSystemTwo{},
		three: &SubSystemThree{},
		four:  &SubSystemFour{},
	}
}

//MethodA 方法组A
func (f *Facade) MethodA() {
	fmt.Println("方法组A()----")
	f.one.MethodOne()
	f.two.MethodTwo()
	f.four.MethodFour()
}

//MethodB 方法组B
func (f *Facade) MethodB() {
	fmt.Println("方法组B()----")
	f.two.MethodTwo()
	f.three.MethodThree()
}
