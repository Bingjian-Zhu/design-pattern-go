package simplefactoryb

import "testing"

func TestOperator(t *testing.T) {
	oper := NewOperation("+")
	strResult := oper.GetResult(1, 2)
	if strResult != 3 {
		t.Fatal("TestOperator test fail")
	}
	oper = NewOperation("*")
	strResult = oper.GetResult(2, 2)
	if strResult != 4 {
		t.Fatal("TestOperator test fail")
	}
}
