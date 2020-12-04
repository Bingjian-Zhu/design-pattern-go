package proxy

func ExampleProxy() {
	proxy := new(Proxy)
	proxy.Request()
	// OutPut:
	// 真实的请求
}
