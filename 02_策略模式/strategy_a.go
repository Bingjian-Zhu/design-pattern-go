package strategya

//CashSuper 现金收费接口
type CashSuper interface {
	AcceptCash(money float64) float64
}

//CashNormal 正常收费
type CashNormal struct {
}

//AcceptCash 正常收费，原价返回
func (*CashNormal) AcceptCash(money float64) float64 {
	return money
}
