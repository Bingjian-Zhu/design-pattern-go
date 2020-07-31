package proxy

import "fmt"

//Pursuit 追求者
type Pursuit struct {
	mm *SchoolGirl
}

//GiveDolls 送洋娃娃
func (p *Pursuit) GiveDolls() {
	fmt.Printf("%s 送你洋娃娃\n", p.mm.Name)
}

//GiveFlowers 送鲜花
func (p *Pursuit) GiveFlowers() {
	fmt.Printf("%s 送你鲜花\n", p.mm.Name)
}

//GiveChocolate 送巧克力
func (p *Pursuit) GiveChocolate() {
	fmt.Printf("%s 送你巧克力\n", p.mm.Name)
}

//SchoolGirl 被追求者
type SchoolGirl struct {
	Name string
}
