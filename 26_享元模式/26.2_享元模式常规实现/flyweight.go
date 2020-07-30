package flyweight

import "fmt"

//Flyweight 接口
type Flyweight interface {
	Operation(extrinsicstate int)
}

//ConcreteFlyweight 具体对象
type ConcreteFlyweight struct {
}

//Operation 操作
func (*ConcreteFlyweight) Operation(extrinsicstate int) {
	fmt.Printf("具体Flyweight:%d\n", extrinsicstate)
}

//UnsharedConcreteFlyweight 非共享对象
type UnsharedConcreteFlyweight struct {
}

//Operation 操作
func (*UnsharedConcreteFlyweight) Operation(extrinsicstate int) {
	fmt.Printf("不共享的具体Flyweight:%d\n", extrinsicstate)
}

//FlyweightFactory 共享对象工厂
type FlyweightFactory struct {
	flyweights map[string]*ConcreteFlyweight
}

//NewFlyweightFactory 对象工厂
func NewFlyweightFactory() *FlyweightFactory {
	f := &FlyweightFactory{
		flyweights: make(map[string]*ConcreteFlyweight),
	}
	f.flyweights["x"] = &ConcreteFlyweight{}
	f.flyweights["y"] = &ConcreteFlyweight{}
	f.flyweights["z"] = &ConcreteFlyweight{}
	return f
}

//GetFlyweight 获取共享对象
func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	return Flyweight(f.flyweights[key])
}
