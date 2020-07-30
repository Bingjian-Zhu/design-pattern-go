package interpreter

func ExampleInterpreter() {
	context := &Context{}
	list := make([]AbstractExpression, 0, 4)
	list = append(list, &TerminalExpression{})
	list = append(list, &NonterminalExpression{})
	list = append(list, &TerminalExpression{})
	list = append(list, &TerminalExpression{})

	for _, exp := range list {
		exp.Interpret(context)
	}

	// OutPut:
	// 终端解释器
	// 非终端解释器
	// 终端解释器
	// 终端解释器

}
