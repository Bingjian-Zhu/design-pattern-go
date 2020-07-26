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
