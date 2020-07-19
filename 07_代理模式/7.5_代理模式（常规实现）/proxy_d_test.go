package proxyd

func ExampleProxy_d() {
	proxy := &Proxy{}
	proxy.Request()
	// OutPut:
	// 真实的请求
}
