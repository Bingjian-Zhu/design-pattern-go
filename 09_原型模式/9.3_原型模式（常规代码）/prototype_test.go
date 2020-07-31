package prototype

import "testing"

var manager *ConcretePrototype

type Type1 struct {
	name string
}

func (t *Type1) Clone() Prototype {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Prototype {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
}

func init() {
	manager = NewConcretePrototype()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}
