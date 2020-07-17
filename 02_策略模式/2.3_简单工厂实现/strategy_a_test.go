package strategya

import "testing"

func TestStrategy_a(t *testing.T) {
	factory := &CashFactory{}
	var cs CashSuper
	cs = factory.CreateCashAccept("正常收费")
	res := cs.AcceptCash(300)
	if res != 300 {
		t.Fatal("CashNormal test fail")
	}
	cs = factory.CreateCashAccept("打八折")
	res = cs.AcceptCash(300)
	if res != 240 {
		t.Fatal("CashRebate test fail")
	}
	cs = factory.CreateCashAccept("满300返100")
	res = cs.AcceptCash(300)
	if res != 200 {
		t.Fatal("CashReturn test fail")
	}
}
