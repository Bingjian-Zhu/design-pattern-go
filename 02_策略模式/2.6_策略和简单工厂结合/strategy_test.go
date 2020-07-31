package strategy

import "testing"

func TestStrategy(t *testing.T) {
	var cc *CashContext
	cc = NewCashContext("正常收费")
	res := cc.GetResult(300)
	if res != 300 {
		t.Fatal("CashNormal test fail")
	}
	cc = NewCashContext("打八折")
	res = cc.GetResult(300)
	if res != 240 {
		t.Fatal("CashRebate test fail")
	}
	cc = NewCashContext("满300返100")
	res = cc.GetResult(300)
	if res != 200 {
		t.Fatal("CashReturn test fail")
	}
}
