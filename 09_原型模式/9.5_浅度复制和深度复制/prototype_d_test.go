package prototyped

import (
	"fmt"
	"testing"
)

var manager *ConcretePrototype

//Type1 浅度复制，Type3为引用
type Type1 struct {
	name string
	*Type3
}

func (t *Type1) Clone() Prototype {
	tc := *t
	return &tc
}

//Type2 深度复制，Type3为值
type Type2 struct {
	name string
	Type3
}

func (t *Type2) Clone() Prototype {
	tc := *t
	return &tc
}

type Type3 struct {
	id string
}

//TestClone1 浅度复制测试
func TestClone1(t *testing.T) {
	manager = NewConcretePrototype()

	a := &Type1{
		name:  "type1",
		Type3: &Type3{id: "t3"},
	}
	manager.Set("t1", a)
	t1 := manager.Get("t1")
	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! 浅度复制出错")
	}
	if &t1 == &t2 {
		t.Fatal("error! 浅度复制出错")
	}

	type1 := t1.(*Type1)
	type2 := t2.(*Type1)

	type1.name = "t1_name"
	if type1.name == type2.name {
		t.Fatal("error! 浅度复制出错")
	}
	type1.Type3.id = "t1_id"
	if type1.Type3.id != type2.Type3.id {
		fmt.Println(type1.Type3.id)
		fmt.Println(type2.Type3.id)
		t.Fatal("error! 浅度复制出错！t3为引用，type1.t3.id == type2.t3.id")
	}
}

//TestClone2 深度复制测试
func TestClone2(t *testing.T) {
	manager = NewConcretePrototype()

	a := &Type2{
		name:  "type2",
		Type3: Type3{id: "t3"}, //此处改成传值
	}
	manager.Set("t1", a)
	t1 := manager.Get("t1")
	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! 深度复制出错")
	}
	if &t1 == &t2 {
		t.Fatal("error! 深度复制出错")
	}

	type1 := t1.(*Type2)
	type2 := t2.(*Type2)

	type1.name = "t1_name"
	if type1.name == type2.name {
		t.Fatal("error! 深度复制出错")
	}
	type1.Type3.id = "t1_id"
	if type1.Type3.id == type2.Type3.id {
		fmt.Println(type1.Type3.id)
		fmt.Println(type2.Type3.id)
		t.Fatal("error! 深度复制出错！t3为值传递，type1.t3.id != type2.t3.id")
	}
}
