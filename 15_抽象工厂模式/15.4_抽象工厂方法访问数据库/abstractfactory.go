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
	return &model.User{}
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
	return &model.User{}
}

//IDepartment 部门接口
type IDepartment interface {
	Insert(user *model.Department)
	GetDepartment(id int) *model.Department
}

//SqlserverDepartment Sqlserver数据库
type SqlserverDepartment struct {
}

//Insert 插入部门
func (*SqlserverDepartment) Insert(user *model.Department) {
	fmt.Println("在Sqlserver中给Department表增加一条记录")
}

//GetDepartment 获取部门
func (*SqlserverDepartment) GetDepartment(id int) *model.Department {
	fmt.Println("在Sqlserver中根据ID得到Department表一条记录")
	return &model.Department{}
}

//AccessDepartment Access
type AccessDepartment struct {
}

//Insert 插入部门
func (*AccessDepartment) Insert(user *model.Department) {
	fmt.Println("在Access中给Department表增加一条记录")
}

//GetDepartment 获取部门
func (*AccessDepartment) GetDepartment(id int) *model.Department {
	fmt.Println("在Access中根据ID得到Department表一条记录")
	return &model.Department{}
}

//IFactory IUser工厂接口
type IFactory interface {
	CreateUser() IUser
	CreateDepartment() IDepartment
}

type SqlserverFactory struct {
}

func (*SqlserverFactory) CreateUser() IUser {
	return &SqlserverUser{}
}

func (*SqlserverFactory) CreateDepartment() IDepartment {
	return &SqlserverDepartment{}
}

type AccessFactory struct {
}

func (*AccessFactory) CreateUser() IUser {
	return &AccessUser{}
}

func (*AccessFactory) CreateDepartment() IDepartment {
	return &AccessDepartment{}
}
