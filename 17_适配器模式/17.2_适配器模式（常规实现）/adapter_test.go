package adapter

func ExampleAdapter() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	target.Request()

	// OutPut:
	// 特殊请求
}
