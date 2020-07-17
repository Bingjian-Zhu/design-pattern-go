package strategyc

import "testing"

func TestStrategy_c(t *testing.T) {
	var cc *CashContext
	strType := "打八折"
	switch strType {
	case "正常收费":
		cc = NewCashContext(&CashNormal{})
		break
	case "满300返100":
		cc = NewCashContext(NewCashReturn(300, 100))
		break
	case "打八折":
		cc = NewCashContext(NewCashRebate(0.8))
		break
	}
	res := cc.GetResult(300)
	if res != 240 {
		t.Fatal("CashRebate test fail")
	}
}
