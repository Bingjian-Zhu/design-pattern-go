package interpreter

func ExampleInterpreter() {
	context := new(Context)
	list := make([]AbstractExpression, 0, 4)
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonterminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))

	for _, exp := range list {
		exp.Interpret(context)
	}

	// OutPut:
	// 终端解释器
	// 非终端解释器
	// 终端解释器
	// 终端解释器

}
