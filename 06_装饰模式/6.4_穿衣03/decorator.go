package decorator

import "fmt"

//Person 人
type Person interface {
	Show()
}

//ConcretePerson 具体的人
type ConcretePerson struct {
	name string
}

//NewConcretePerson ConcretePerson构造函数
func NewConcretePerson(name string) *ConcretePerson {
	return &ConcretePerson{
		name: name,
	}
}

//Show 显示人名
func (c *ConcretePerson) Show() {
	fmt.Printf("装扮的%s\n", c.name)
}

//Finery 服饰
type Finery struct {
	Person
}

//Decorate 服饰装饰器
func (f *Finery) Decorate(component Person) {
	f.Person = component
}

//Show 服饰展示
func (f *Finery) Show() {
	if f.Person != nil {
		f.Person.Show()
	}
}

//TShirts 大T恤
type TShirts struct {
	Person
}

//Show 大T恤
func (t *TShirts) Show() {
	fmt.Print("大T恤 ")
	t.Person.Show()
}

//BigTrouser 垮裤
type BigTrouser struct {
	Person
}

//Show 垮裤
func (b *BigTrouser) Show() {
	fmt.Print("垮裤 ")
	b.Person.Show()
}

//Sneakers 破球鞋
type Sneakers struct {
	Person
}

//Show 破球鞋
func (s *Sneakers) Show() {
	fmt.Print("破球鞋 ")
	s.Person.Show()
}

//Suit 西装
type Suit struct {
	Person
}

//Show 西装
func (s *Suit) Show() {
	fmt.Print("西装 ")
	s.Person.Show()
}

//Tie 领带
type Tie struct {
	Person
}

//Show 领带
func (t *Tie) Show() {
	fmt.Print("领带 ")
	t.Person.Show()
}

//LeatherShoes 皮鞋
type LeatherShoes struct {
	Person
}

//Show 皮鞋
func (f *LeatherShoes) Show() {
	fmt.Print("皮鞋 ")
	f.Person.Show()
}
