package memento

import "fmt"

//Originator 发起人
type Originator struct {
	State string
}

//CreateMemento 构造函数
func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		State: o.State,
	}
}

//SetMemento 设置备忘录
func (o *Originator) SetMemento(memento *Memento) {
	o.State = memento.State
}

//Show 展示
func (o *Originator) Show() {
	fmt.Println("State=" + o.State)
}

//Memento 备忘录
type Memento struct {
	State string
}

//Caretaker 管理者
type Caretaker struct {
	*Memento
}
