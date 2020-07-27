package composite

import (
	"fmt"
	"strings"
)

//Component 接口
type Component interface {
	Add(c Component)
	Remove(c Component)
	Display(depth int)
}

//Composite 组合
type Composite struct {
	name     string
	children []Component
}

//Add 添加
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

//Remove 移除
func (c *Composite) Remove(component Component) {
	for index, val := range c.children {
		if component == val {
			c.children = append(c.children[0:index], c.children[index+1:]...)
			break
		}
	}
}

//Display 展示
func (c *Composite) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + c.name)
	for _, component := range c.children {
		component.Display(depth + 2)
	}
}

//Leaf 叶子
type Leaf struct {
	name string
}

//Add 添加
func (*Leaf) Add(c Component) {
	fmt.Println("Cannot add to a leaf")
}

//Remove 移除
func (*Leaf) Remove(c Component) {
	fmt.Println("Cannot remove from a leaf")
}

//Display 展示
func (l *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + l.name)
}
