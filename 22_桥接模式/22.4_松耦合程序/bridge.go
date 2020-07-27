package bridge

import "fmt"

//HandsetSoft 手机软件
type HandsetSoft interface {
	Run()
}

//HandsetAddressList 通讯录
type HandsetAddressList interface {
	Run()
}

//HandsetGame 游戏
type HandsetGame interface {
	Run()
}

//HandsetBrandMGame 手机品牌M的游戏
type HandsetBrandMGame struct {
}

//Run 运行
func (*HandsetBrandMGame) Run() {
	fmt.Println("运行M品牌手机游戏")
}

//HandsetBrandNGame 手机品牌N的游戏
type HandsetBrandNGame struct {
}

//Run 运行
func (*HandsetBrandNGame) Run() {
	fmt.Println("运行N品牌手机游戏")
}

//HandsetBrandMAddressList 手机品牌M的通讯录
type HandsetBrandMAddressList struct {
}

//Run 运行
func (*HandsetBrandMAddressList) Run() {
	fmt.Println("运行M品牌手机通讯录")
}

//HandsetBrandNAddressList 手机品牌N的通讯录
type HandsetBrandNAddressList struct {
}

//Run 运行
func (*HandsetBrandNAddressList) Run() {
	fmt.Println("运行N品牌手机通讯录")
}
