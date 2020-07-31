package obserser

import "fmt"

//Observer 观察者接口
type Observer interface {
	Update()
}

//Subject 通知者接口
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
	SubjectState() string
}

//Secretary 前台秘书类
type Secretary struct {
	observers []Observer
	Action    string
}

//Attach 增加
func (s *Secretary) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

//Detach 移除
func (s *Secretary) Detach(observer Observer) {
	for index, val := range s.observers {
		if val == observer {
			s.observers = append(s.observers[:index], s.observers[index+1:]...)
			break
		}
	}
}

//Notify 通知
func (s *Secretary) Notify() {
	for _, val := range s.observers {
		val.Update()
	}
}

//SubjectState 消息
func (s *Secretary) SubjectState() string {
	return s.Action
}

//Boss 老板
type Boss struct {
	observers []Observer
	Action    string
}

//Attach 增加
func (s *Boss) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

//Detach 移除
func (s *Boss) Detach(observer Observer) {
	for index, val := range s.observers {
		if val == observer {
			s.observers = append(s.observers[:index], s.observers[index+1:]...)
			break
		}
	}
}

//Notify 通知
func (s *Boss) Notify() {
	for _, val := range s.observers {
		val.Update()
	}
}

//SubjectState 消息
func (s *Boss) SubjectState() string {
	return s.Action
}

//StockObserver 看股票的同事
type StockObserver struct {
	name string
	Subject
}

//NewStockObserver StockObserver构造函数
func NewStockObserver(name string, s Subject) *StockObserver {
	return &StockObserver{
		name:    name,
		Subject: s,
	}
}

//Update 更新状态
func (s *StockObserver) Update() {
	fmt.Printf("%s %s 关闭股票行情，继续工作！\n", s.Subject.SubjectState(), s.name)
}

//NBAObserver 看NBA的同事
type NBAObserver struct {
	name string
	Subject
}

//NewNBAObserver NBAObserver
func NewNBAObserver(name string, s Subject) *NBAObserver {
	return &NBAObserver{
		name:    name,
		Subject: s,
	}
}

//Update 更新状态
func (s *NBAObserver) Update() {
	fmt.Printf("%s %s 关闭NBA直播，继续工作！\n", s.Subject.SubjectState(), s.name)
}
