package strategya

import "testing"

func TestStrategy_a(t *testing.T) {
	var cs CashSuper
	cs = &CashNormal{}
	res := cs.AcceptCash(100)
	if res != 100 {
		t.Fatal("CashNormal test fail")
	}
	cs = NewCashRebate(0.8)
	res = cs.AcceptCash(100)
	if res != 80 {
		t.Fatal("CashRebate test fail")
	}
	cs = NewCashReturn(300, 100)
	res = cs.AcceptCash(300)
	if res != 200 {
		t.Fatal("CashReturn test fail")
	}
}
