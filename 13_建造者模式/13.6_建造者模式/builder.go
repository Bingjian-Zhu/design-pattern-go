package builder

import "fmt"

//Product 产品
type Product struct {
	parts []string
}

//Add 添加产品
func (p *Product) Add(part string) {
	p.parts = append(p.parts, part)
}

//Show 展示产品
func (p *Product) Show() {
	fmt.Println("产品 创建 ----")
	for _, part := range p.parts {
		fmt.Println(part)
	}
}

//Builder 建造接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() *Product
}

//ConcreteBuilder1 具体产品建造
type ConcreteBuilder1 struct {
	product *Product
}

//NewConcreteBuilder1 ConcreteBuilder1构造函数
func NewConcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{
		product: new(Product),
	}
}

//BuildPartA 建造部件A
func (c *ConcreteBuilder1) BuildPartA() {
	c.product.Add("部件A")
}

//BuildPartB 建造部件B
func (c *ConcreteBuilder1) BuildPartB() {
	c.product.Add("部件B")
}

//GetResult 返回产品
func (c *ConcreteBuilder1) GetResult() *Product {
	return c.product
}

//ConcreteBuilder2 具体产品建造
type ConcreteBuilder2 struct {
	product *Product
}

//NewConcreteBuilder2 ConcreteBuilder2构造函数
func NewConcreteBuilder2() *ConcreteBuilder2 {
	return &ConcreteBuilder2{
		product: new(Product),
	}
}

//BuildPartA 建造部件X
func (c *ConcreteBuilder2) BuildPartA() {
	c.product.Add("部件X")
}

//BuildPartB 建造部件Y
func (c *ConcreteBuilder2) BuildPartB() {
	c.product.Add("部件Y")
}

//GetResult 返回产品
func (c *ConcreteBuilder2) GetResult() *Product {
	return c.product
}

//Director 指挥者
type Director struct {
}

//Construct 指挥建造过程
func (*Director) Construct(builder Builder) {
	builder.BuildPartA()
	builder.BuildPartB()
}
