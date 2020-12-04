package mediator

func ExampleMediator() {
	m := new(ConcreteMediator)

	c1 := NewConcreteColleague1(m)
	c2 := NewConcreteColleague2(m)

	m.Colleague1 = c1
	m.Colleague2 = c2

	c1.Send("吃过饭了吗?")
	c2.Send("没有呢，你打算请客？")

	// OutPut:
	// 同事2得到信息:吃过饭了吗?
	// 同事1得到信息:没有呢，你打算请客？
}
