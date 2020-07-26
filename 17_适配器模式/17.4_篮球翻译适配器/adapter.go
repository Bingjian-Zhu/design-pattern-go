package adapter

import "fmt"

//Player 篮球运动员接口
type Player interface {
	Attack()
	Defense()
}

//Forwards 前锋
type Forwards struct {
	Name string
}

//Attack 进攻
func (a *Forwards) Attack() {
	fmt.Printf("前锋 %s 进攻\n", a.Name)
}

//Defense 防守
func (a *Forwards) Defense() {
	fmt.Printf("前锋 %s 防守\n", a.Name)
}

//Center 中锋
type Center struct {
	Name string
}

//Attack 进攻
func (c *Center) Attack() {
	fmt.Printf("中锋 %s 进攻\n", c.Name)
}

//Defense 防守
func (c *Center) Defense() {
	fmt.Printf("中锋 %s 防守\n", c.Name)
}

//Guards 后卫
type Guards struct {
	Name string
}

//Attack 进攻
func (g *Guards) Attack() {
	fmt.Printf("后卫 %s 进攻\n", g.Name)
}

//Defense 防守
func (g *Guards) Defense() {
	fmt.Printf("后卫 %s 防守\n", g.Name)
}

//ForeignCenter 外籍中锋
type ForeignCenter struct {
	Name string
}

//Attack 进攻
func (f *ForeignCenter) 进攻() {
	fmt.Printf("外籍中锋 %s 进攻\n", f.Name)
}

//Defense 防守
func (f *ForeignCenter) 防守() {
	fmt.Printf("外籍中锋 %s 防守\n", f.Name)
}

//NewTranslator 构造函数
func NewTranslator(name string) *Translator {
	return &Translator{
		wjzf: &ForeignCenter{
			Name: name,
		},
	}
}

//Translator 翻译者
type Translator struct {
	wjzf *ForeignCenter
}

//Attack 进攻
func (t *Translator) Attack() {
	t.wjzf.进攻()
}

//Defense 防守
func (t *Translator) Defense() {
	t.wjzf.防守()
}
