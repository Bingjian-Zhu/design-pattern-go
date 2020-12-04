package memento

func ExampleMemento() {
	//大战Boss前
	lixiaoyao := new(GameRole)
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()

	//保存进度
	stateAdmin := new(RoleStateCaretaker)
	stateAdmin.Memento = lixiaoyao.SaveState()

	//大战Boss时，损耗严重
	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()

	//恢复之前状态
	lixiaoyao.RecoveryState(stateAdmin.Memento)

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
