package proxy

import "fmt"

//Proxy 代理
type Proxy struct {
	mm *SchoolGirl
}

//GiveDolls 送洋娃娃
func (p *Proxy) GiveDolls() {
	fmt.Printf("%s 送你洋娃娃\n", p.mm.Name)
}

//GiveFlowers 送鲜花
func (p *Proxy) GiveFlowers() {
	fmt.Printf("%s 送你鲜花\n", p.mm.Name)
}

//GiveChocolate 送巧克力
func (p *Proxy) GiveChocolate() {
	fmt.Printf("%s 送你巧克力\n", p.mm.Name)
}

//SchoolGirl 被追求者
type SchoolGirl struct {
	Name string
}
