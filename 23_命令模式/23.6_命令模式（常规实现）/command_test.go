package command

func ExampleCommand() {
	r := &Receiver{}
	c := NewConcreteCommand(r)
	i := &Invoker{}

	// Set and execute command
	i.SetCommand(c)
	i.ExecuteCommand()

	// OutPut:
	// 执行请求！
}
