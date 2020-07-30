package flyweight

import "fmt"

func ExampleFlyweight() {
	f := &WebSiteFactory{
		flyweights: make(map[string]*ConcreteWebSite),
	}

	fx := f.GetWebSiteCategory("产品展示")
	fx.Use()

	fy := f.GetWebSiteCategory("产品展示")
	fy.Use()

	fz := f.GetWebSiteCategory("产品展示")
	fz.Use()

	fl := f.GetWebSiteCategory("博客")
	fl.Use()

	fm := f.GetWebSiteCategory("博客")
	fm.Use()

	fn := f.GetWebSiteCategory("博客")
	fn.Use()

	fmt.Printf("网站分类总数为 %d", f.GetWebSiteCount())

	// OutPut:
	// 网站分类：产品展示
	// 网站分类：产品展示
	// 网站分类：产品展示
	// 网站分类：博客
	// 网站分类：博客
	// 网站分类：博客
	// 网站分类总数为 2

}
