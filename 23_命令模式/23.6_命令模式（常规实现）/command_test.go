package command

func ExampleCommand() {
	r := new(Receiver)
	c := NewConcreteCommand(r)
	i := new(Invoker)

	// Set and execute command
	i.SetCommand(c)
	i.ExecuteCommand()

	// OutPut:
	// 执行请求！
}
