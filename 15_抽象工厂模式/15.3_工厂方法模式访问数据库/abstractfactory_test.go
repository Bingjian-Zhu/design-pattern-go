package abstractfactory

import "design-pattern-go/15_抽象工厂模式/model"

func ExampleAbstractfactory() {
	user := new(model.User)

	factory := new(SqlserverFactory)
	su := factory.CreateUser()
	su.Insert(user)
	su.GetUser(1)

	// OutPut:
	// 在Sqlserver中给User表增加一条记录
	// 在Sqlserver中根据ID得到User表一条记录
}
