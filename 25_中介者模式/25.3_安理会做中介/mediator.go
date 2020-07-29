package mediator

import "fmt"

//UnitedNations 联合国机构
type UnitedNations interface {
	Declare(message string, colleague Country)
}

//UnitedNationsSecurityCouncil 联合国安全理事会
type UnitedNationsSecurityCouncil struct {
	Colleague1 *USA
	Colleague2 *Iraq
}

//Declare 声明
func (u *UnitedNationsSecurityCouncil) Declare(message string, colleague Country) {
	if colleague == u.Colleague1 {
		u.Colleague2.GetMessage(message)
	} else {
		u.Colleague1.GetMessage(message)
	}
}

//Country 国家
type Country interface {
	Declare(message string)
	GetMessage(message string)
}

//USA 美国
type USA struct {
	mediator UnitedNations
}

//NewUSA 美国构造函数
func NewUSA(mediator UnitedNations) *USA {
	return &USA{
		mediator: mediator,
	}
}

//Declare 声明
func (u *USA) Declare(message string) {
	u.mediator.Declare(message, u)
}

//GetMessage 获得消息
func (u *USA) GetMessage(message string) {
	fmt.Println("美国获得对方信息：" + message)
}

//Iraq 伊拉克
type Iraq struct {
	mediator UnitedNations
}

//NewIraq 伊拉克构造函数
func NewIraq(mediator UnitedNations) *Iraq {
	return &Iraq{
		mediator: mediator,
	}
}

//Declare 声明
func (i *Iraq) Declare(message string) {
	i.mediator.Declare(message, i)
}

//GetMessage 获得消息
func (i *Iraq) GetMessage(message string) {
	fmt.Println("伊拉克获得对方信息：" + message)
}
