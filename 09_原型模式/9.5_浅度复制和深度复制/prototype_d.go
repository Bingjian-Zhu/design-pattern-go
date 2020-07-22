package prototyped

//Prototype 原型
type Prototype interface {
	Clone() Prototype
}

//ConcretePrototype 具体对象
type ConcretePrototype struct {
	pMap map[string]Prototype
}

//NewConcretePrototype ConcretePrototype构造函数
func NewConcretePrototype() *ConcretePrototype {
	return &ConcretePrototype{
		pMap: make(map[string]Prototype),
	}
}

//Get 获取对象
func (c *ConcretePrototype) Get(id string) Prototype {
	return c.pMap[id]
}

//Set 设置对象
func (c *ConcretePrototype) Set(id string, cloneable Prototype) {
	c.pMap[id] = cloneable
}
