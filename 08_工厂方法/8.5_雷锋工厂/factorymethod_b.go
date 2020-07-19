package factorymethodb

import "fmt"

//LeiFeng 雷锋
type LeiFeng struct {
}

//Sweep 扫地
func (*LeiFeng) Sweep() {
	fmt.Println("扫地")
}

//Wash 洗衣
func (*LeiFeng) Wash() {
	fmt.Println("洗衣")
}

//BuyRice 买米
func BuyRice() {
	fmt.Println("买米")
}

//Undergraduate 学雷锋的大学生
type Undergraduate struct {
	LeiFeng
}

//Volunteer 社区志愿者
type Volunteer struct {
	LeiFeng
}

//SimpleFactory 简单雷锋工厂
type SimpleFactory struct {
}

//CreateLeiFeng 新建实例
func (*SimpleFactory) CreateLeiFeng(strType string) *LeiFeng {
	var result *LeiFeng
	switch strType {
	case "学雷锋的大学生":
		result = &Undergraduate{}
		break
	case "社区志愿者":
		result = &Volunteer{}
		break
	}
	return result
}

//IFactory 雷锋工厂
type IFactory interface {
	CreateLeiFeng() LeiFeng
}

//UndergraduateFactory 学雷锋的大学生工厂
type UndergraduateFactory struct {
}

//CreateLeiFeng 新建学雷锋的大学生工厂实例
func (*UndergraduateFactory) CreateLeiFeng() LeiFeng {
	return &Undergraduate{}
}

//VolunteerFactory 社区志愿者工厂
type VolunteerFactory struct {
}

//CreateLeiFeng 新建社区志愿者工厂实例
func (*VolunteerFactory) CreateLeiFeng() LeiFeng {
	return &Volunteer{}
}
