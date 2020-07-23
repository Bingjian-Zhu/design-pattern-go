package observer

import "fmt"

//Subject 通知者
type Subject struct {
	observers []Observer
	context   string
}

//NewSubject Subject构造函数
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

//Attach 添加
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

//notify 通知
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

//UpdateContext 发出通知
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

//Observer 观察者接口
type Observer interface {
	Update(*Subject)
}

//Reader 订阅者
type Reader struct {
	name string
}

//NewReader Reader构造函数
func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

//Update 订阅者更新状态
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
