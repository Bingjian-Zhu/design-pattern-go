package factorymethod

import "log"

//Operation 运算结构体
type Operation interface {
	GetResult(numberA, numberB float64) float64
}

//OperationAdd 加法运算
type OperationAdd struct {
}

//GetResult 获取加法结果
func (*OperationAdd) GetResult(numberA, numberB float64) float64 {
	return numberA + numberB
}

//OperationSub 减法运算
type OperationSub struct {
}

//GetResult 获取减法结果
func (*OperationSub) GetResult(numberA, numberB float64) float64 {
	return numberA - numberB
}

//OperationMul 乘法运算
type OperationMul struct {
}

//GetResult 获取乘法结果
func (*OperationMul) GetResult(numberA, numberB float64) float64 {
	return numberA * numberB
}

//OperationDiv 除法运算
type OperationDiv struct {
}

//GetResult 获取除法结果
func (*OperationDiv) GetResult(numberA, numberB float64) float64 {
	if numberB == 0 {
		log.Fatal("除数不能为0")
	}
	return numberA / numberB
}

//IFactory 工厂方法
type IFactory interface {
	CreateOperation()
}

//AddFactory 专门负责生产“+”的工厂
type AddFactory struct {
}

//CreateOperation 创建+法实例
func (*AddFactory) CreateOperation() Operation {
	return &OperationAdd{}
}

//SubFactory 专门负责生产“-”的工厂
type SubFactory struct {
}

//CreateOperation 创建-法实例
func (*SubFactory) CreateOperation() Operation {
	return &OperationSub{}
}

//MulFactory 专门负责生产“*”的工厂
type MulFactory struct {
}

//CreateOperation 创建*法实例
func (*MulFactory) CreateOperation() Operation {
	return &OperationMul{}
}

//DivFactory 专门负责生产“/”的工厂
type DivFactory struct {
}

//CreateOperation 创建*法实例
func (*DivFactory) CreateOperation() Operation {
	return &OperationDiv{}
}
