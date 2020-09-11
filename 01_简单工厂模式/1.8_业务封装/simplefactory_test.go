package simplefactory

import "testing"

func TestOperator(t *testing.T) {
	oper := new(Operation)
	strResult := oper.GetResult(1, 2, "+")
	if strResult != 3 {
		t.Fatal("TestOperator test fail")
	}
}
