package interpreter

import "fmt"

//Context 上下文
type Context struct {
	Input  string
	Output string
}

//AbstractExpression 解析器接口
type AbstractExpression interface {
	Interpret(context *Context)
}

//TerminalExpression 终端解释器
type TerminalExpression struct {
}

//Interpret 解析操作
func (*TerminalExpression) Interpret(context *Context) {
	fmt.Println("终端解释器")
}

//NonterminalExpression 非终端解释器
type NonterminalExpression struct {
}

//Interpret 解析操作
func (*NonterminalExpression) Interpret(context *Context) {
	fmt.Println("非终端解释器")
}
