package simplefactory

//Operation 运算结构体
type Operation struct {
}

//GetResult 获取运算结果
func (*Operation) GetResult(numberA, numberB float64, operate string) float64 {
	var result float64 = 0
	switch operate {
	case "+":
		result = numberA + numberB
		break
	case "-":
		result = numberA - numberB
		break
	case "*":
		result = numberA * numberB
		break
	case "/":
		result = numberA / numberB
		break
	}
	return result
}
