package memento

func ExampleMemento() {
	//大战Boss前
	lixiaoyao := new(GameRole)
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()

	//保存进度
	backup := new(GameRole)
	backup.Vitality = lixiaoyao.Vitality
	backup.Attack = lixiaoyao.Attack
	backup.Defense = lixiaoyao.Defense

	//大战Boss时，损耗严重
	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()

	//恢复之前状态
	lixiaoyao.Vitality = backup.Vitality
	lixiaoyao.Attack = backup.Attack
	lixiaoyao.Defense = backup.Defense

	lixiaoyao.StateDisplay()

	// OutPut:
	// 角色当前状态：
	// 体力：100
	// 攻击力：100
	// 防御力：100
	//
	// 角色当前状态：
	// 体力：0
	// 攻击力：0
	// 防御力：0
	//
	// 角色当前状态：
	// 体力：100
	// 攻击力：100
	// 防御力：100
}
