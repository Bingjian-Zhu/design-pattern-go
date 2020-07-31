package obserser

import "fmt"

//Secretary 前台秘书类
type Secretary struct {
	observers []*StockObserver
	Action    string
}

//Attach 增加
func (s *Secretary) Attach(stockObserver *StockObserver) {
	s.observers = append(s.observers, stockObserver)
}

//Detach 移除
func (s *Secretary) Detach(stockObserver *StockObserver) {
	for index, val := range s.observers {
		if val == stockObserver {
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

//StockObserver 看股票的同事
type StockObserver struct {
	name string
	*Secretary
}

//NewStockObserver StockObserver构造函数
func NewStockObserver(name string, s *Secretary) *StockObserver {
	return &StockObserver{
		name:      name,
		Secretary: s,
	}
}

//Update 更新状态
func (s *StockObserver) Update() {
	fmt.Printf("%s %s 关闭股票行情，继续工作！\n", s.Secretary.Action, s.name)
}
