package flyweight

import "fmt"

//WebSite 网站
type WebSite interface {
	Use()
}

//ConcreteWebSite 具体的网站
type ConcreteWebSite struct {
	Name string
}

//Use 使用
func (c *ConcreteWebSite) Use() {
	fmt.Println("网站分类：" + c.Name)
}

//WebSiteFactory 网站工厂
type WebSiteFactory struct {
	flyweights map[string]*ConcreteWebSite
}

//GetWebSiteCategory 获得网站分类
func (w *WebSiteFactory) GetWebSiteCategory(key string) WebSite {
	c := &ConcreteWebSite{}
	c.Name = key
	w.flyweights[key] = c
	return WebSite(w.flyweights[key])
}

//GetWebSiteCount 获得网站分类总数
func (w *WebSiteFactory) GetWebSiteCount() int {
	return len(w.flyweights)
}
