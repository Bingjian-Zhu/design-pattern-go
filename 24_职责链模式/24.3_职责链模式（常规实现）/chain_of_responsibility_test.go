package chain

func ExampleChain() {
	h1 := NewConcreteHandler1()
	h2 := NewConcreteHandler2()
	h3 := NewConcreteHandler3()
	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)

	requests := []int{2, 5, 14, 22, 18, 3, 27, 20}
	for _, val := range requests {
		h1.HandleRequest(val)
	}

	// OutPut:
	// ConcreteHandler1  处理请求  2
	// ConcreteHandler1  处理请求  5
	// ConcreteHandler2  处理请求  14
	// ConcreteHandler3  处理请求  22
	// ConcreteHandler2  处理请求  18
	// ConcreteHandler1  处理请求  3
	// ConcreteHandler3  处理请求  27
	// ConcreteHandler3  处理请求  20
}
