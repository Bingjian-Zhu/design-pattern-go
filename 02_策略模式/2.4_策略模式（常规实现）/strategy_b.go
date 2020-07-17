package strategyb

import "fmt"

//Strategy 抽象算法类
type Strategy interface {
	AlgorithmInterface()
}

//ConcreteStrategyA 具体算法A
type ConcreteStrategyA struct {
}

//AlgorithmInterface 实现算法A
func (*ConcreteStrategyA) AlgorithmInterface() {
	fmt.Println("算法A实现")
}

//ConcreteStrategyB 具体算法B
type ConcreteStrategyB struct {
}

//AlgorithmInterface 实现算法B
func (*ConcreteStrategyB) AlgorithmInterface() {
	fmt.Println("算法B实现")
}

//ConcreteStrategyC 具体算法C
type ConcreteStrategyC struct {
}

//AlgorithmInterface 实现算法C
func (*ConcreteStrategyC) AlgorithmInterface() {
	fmt.Println("算法C实现")
}

//Context 上下文
type Context struct {
	strategy Strategy
}

//NewContext 上下文构造函数
func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

//ContextInterface 上下文接口
func (context *Context) ContextInterface() {
	context.strategy.AlgorithmInterface()
}
