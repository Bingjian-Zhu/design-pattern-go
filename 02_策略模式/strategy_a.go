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

//CashRebate 打折收费
type CashRebate struct {
	moneyRebate float64
}

//NewCashRebate 构造函数，返回CashRebate
func NewCashRebate(moneyRebate float64) CashSuper {
	return &CashRebate{
		moneyRebate: moneyRebate,
	}
}

//AcceptCash 打折收费，乘以折扣
func (rebate *CashRebate) AcceptCash(money float64) float64 {
	return money * rebate.moneyRebate
}

//CashReturn 返利收费
type CashReturn struct {
	moneyCondition, moneyReturn float64
}

//NewCashReturn 构造函数，返回CashReturn
func NewCashReturn(moneyCondition, moneyReturn float64) CashSuper {
	return &CashReturn{
		moneyCondition: moneyCondition,
		moneyReturn:    moneyReturn,
	}
}

//AcceptCash 返利收费
func (re *CashReturn) AcceptCash(money float64) float64 {
	if money >= re.moneyCondition {
		money -= re.moneyReturn
	}
	return money
}
