package factorymethod

import "testing"

func TestFactorymethoda(t *testing.T) {
	operFactory := new(AddFactory)
	oper := operFactory.CreateOperation()
	res := oper.GetResult(1, 2)
	if res != 3 {
		t.Fatal("AddFactory test fail")
	}
}
