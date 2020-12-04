package abstractfactory

import (
	"design-pattern-go/15_抽象工厂模式/model"
	"fmt"
)

//IUser 用户接口
type IUser interface {
	Insert(user *model.User)
	GetUser(id int) *model.User
}

//SqlserverUser Sqlserver数据库
type SqlserverUser struct {
}

//Insert 插入用户
func (*SqlserverUser) Insert(user *model.User) {
	fmt.Println("在Sqlserver中给User表增加一条记录")
}

//GetUser 获取用户
func (*SqlserverUser) GetUser(id int) *model.User {
	fmt.Println("在Sqlserver中根据ID得到User表一条记录")
	return new(model.User)
}

//AccessUser Access
type AccessUser struct {
}

//Insert 插入用户
func (*AccessUser) Insert(user *model.User) {
	fmt.Println("在Access中给User表增加一条记录")
}

//GetUser 获取用户
func (*AccessUser) GetUser(id int) *model.User {
	fmt.Println("在Access中根据ID得到User表一条记录")
	return new(model.User)
}

//IFactory IUser工厂接口
type IFactory interface {
	CreateUser() IUser
}

type SqlserverFactory struct {
}

func (*SqlserverFactory) CreateUser() IUser {
	return new(SqlserverUser)
}

type AccessFactory struct {
}

func (*AccessFactory) CreateUser() IUser {
	return new(AccessUser)
}
