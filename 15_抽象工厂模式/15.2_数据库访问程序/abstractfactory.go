package abstractfactory

import (
	"design-pattern-go/15_抽象工厂模式/model"
	"fmt"
)

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
