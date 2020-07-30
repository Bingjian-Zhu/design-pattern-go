package visitor

import (
	"fmt"
	"reflect"
)

//Person 人类
type Person interface {
	GetConclusion()
}

//Man 男人
type Man struct {
	Action string
}

//GetConclusion 得到结论或反应
func (m *Man) GetConclusion() {
	if m.Action == "成功" {
		fmt.Printf("%s%s时，背后多半有一个伟大的女人。\n", reflect.TypeOf(m).Elem().Name(), m.Action)
	} else if m.Action == "失败" {
		fmt.Printf("%s%s时，闷头喝酒，谁也不用劝。\n", reflect.TypeOf(m).Elem().Name(), m.Action)
	} else if m.Action == "恋爱" {
		fmt.Printf("%s%s时，凡事不懂也要装懂。\n", reflect.TypeOf(m).Elem().Name(), m.Action)
	}
}

//Woman 女人
type Woman struct {
	Action string
}

//GetConclusion 得到结论或反应
func (w *Woman) GetConclusion() {
	if w.Action == "成功" {
		fmt.Printf("%s%s时，背后大多有一个不成功的男人。\n", reflect.TypeOf(w).Elem().Name(), w.Action)
	} else if w.Action == "失败" {
		fmt.Printf("%s%s时，眼泪汪汪，谁也劝不了。\n", reflect.TypeOf(w).Elem().Name(), w.Action)
	} else if w.Action == "恋爱" {
		fmt.Printf("%s%s时，遇事懂也装作不懂。\n", reflect.TypeOf(w).Elem().Name(), w.Action)
	}
}
