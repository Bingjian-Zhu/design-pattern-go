package simplefactory

import "log"

//Operation 运算结构体
type Operation interface {
	GetResult(numberA, numberB float64) float64
}

//NewOperation 返回运算示例
func NewOperation(operate string) Operation {
	var oper Operation
	switch operate {
	case "+":
		oper = new(OperationAdd)
		break
	case "-":
		oper = new(OperationSub)
		break
	case "*":
		oper = new(OperationMul)
		break
	case "/":
		oper = new(OperationDiv)
		break
	}
	return oper
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
