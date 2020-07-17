package decoratora

import "fmt"

//Person 人
type Person struct {
	name string
}

//NewPerson Person构造函数
func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

//WearTShirts 大T恤
func (*Person) WearTShirts() {
	fmt.Print("大T恤 ")
}

//WearBigTrouser 垮裤
func (*Person) WearBigTrouser() {
	fmt.Print("垮裤 ")
}

//WearSneakers 破球鞋
func (*Person) WearSneakers() {
	fmt.Print("破球鞋 ")
}

//WearSuit 西装
func (*Person) WearSuit() {
	fmt.Print("西装 ")
}

//WearTie 领带
func (*Person) WearTie() {
	fmt.Print("领带 ")
}

//WearLeatherShoes 皮鞋
func (*Person) WearLeatherShoes() {
	fmt.Print("皮鞋 ")
}

//Show 显示人名
func (person *Person) Show() {
	fmt.Printf("装扮的%s\n", person.name)
}
