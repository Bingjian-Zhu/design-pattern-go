package decorator

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

//Show 显示人名
func (person *Person) Show() {
	fmt.Printf("装扮的%s\n", person.name)
}

//Finery 服饰
type Finery interface {
	Show()
}

//TShirts 大T恤
type TShirts struct {
}

//Show 大T恤
func (*TShirts) Show() {
	fmt.Print("大T恤 ")
}

//BigTrouser 垮裤
type BigTrouser struct {
}

//Show 垮裤
func (*BigTrouser) Show() {
	fmt.Print("垮裤 ")
}

//Sneakers 破球鞋
type Sneakers struct {
}

//Show 破球鞋
func (*Sneakers) Show() {
	fmt.Print("破球鞋 ")
}

//Suit 西装
type Suit struct {
}

//Show 西装
func (*Suit) Show() {
	fmt.Print("西装 ")
}

//Tie 领带
type Tie struct {
}

//Show 领带
func (*Tie) Show() {
	fmt.Print("领带 ")
}

//LeatherShoes 皮鞋
type LeatherShoes struct {
}

//Show 皮鞋
func (*LeatherShoes) Show() {
	fmt.Print("皮鞋 ")
}
