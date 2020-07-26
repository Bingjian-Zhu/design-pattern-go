package memento

import "fmt"

//GameRole 角色
type GameRole struct {
	//生命力
	Vitality int
	//攻击力
	Attack int
	//防御力
	Defense int
}

//StateDisplay 状态显示
func (g *GameRole) StateDisplay() {
	fmt.Printf("角色当前状态：\n")
	fmt.Printf("体力：%d\n", g.Vitality)
	fmt.Printf("攻击力：%d\n", g.Attack)
	fmt.Printf("防御力：%d\n", g.Defense)
	fmt.Println("")
}

//GetInitState 获得初始状态
func (g *GameRole) GetInitState() {
	g.Vitality = 100
	g.Attack = 100
	g.Defense = 100
}

//Fight 战斗
func (g *GameRole) Fight() {
	g.Vitality = 0
	g.Attack = 0
	g.Defense = 0
}

//SaveState 保存角色状态
func (g *GameRole) SaveState() *RoleStateMemento {
	return &RoleStateMemento{
		Vitality: g.Vitality,
		Attack:   g.Attack,
		Defense:  g.Defense,
	}
}

//RecoveryState 恢复角色状态
func (g *GameRole) RecoveryState(memento *RoleStateMemento) {
	g.Vitality = memento.Vitality
	g.Attack = memento.Attack
	g.Defense = memento.Defense
}

//RoleStateMemento 角色状态存储箱
type RoleStateMemento struct {
	//生命力
	Vitality int
	//攻击力
	Attack int
	//防御力
	Defense int
}

//RoleStateCaretaker 角色状态管理者
type RoleStateCaretaker struct {
	Memento *RoleStateMemento
}
