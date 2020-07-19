package proxyd

import "fmt"

//Subject 接口
type Subject interface {
	Request()
}

//RealSubject 真实的请求
type RealSubject struct {
}

//Request 请求
func (*RealSubject) Request() {
	fmt.Println("真实的请求")
}

//Proxy 代理
type Proxy struct {
	realSubject *RealSubject
}

//Request 代理请求
func (p *Proxy) Request() {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	p.realSubject.Request()
}
