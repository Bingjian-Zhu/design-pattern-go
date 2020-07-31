package facade

import "fmt"

//Stock1 股票1
type Stock1 struct {
}

//Sell 卖出
func (*Stock1) Sell() {
	fmt.Println("股票1卖出")
}

//Buy 买股票
func (*Stock1) Buy() {
	fmt.Println("股票1买入")
}

//Stock2 股票2
type Stock2 struct {
}

//Sell 卖出
func (*Stock2) Sell() {
	fmt.Println("股票2卖出")
}

//Buy 买股票
func (*Stock2) Buy() {
	fmt.Println("股票2买入")
}

//Stock3 股票3
type Stock3 struct {
}

//Sell 卖出
func (*Stock3) Sell() {
	fmt.Println("股票3卖出")
}

//Buy 买股票
func (*Stock3) Buy() {
	fmt.Println("股票3买入")
}

//NationalDebt1 国债1
type NationalDebt1 struct {
}

//Sell 卖国债
func (*NationalDebt1) Sell() {
	fmt.Println("国债1卖出")
}

//Buy 买国债
func (*NationalDebt1) Buy() {
	fmt.Println("国债1买入")
}

//Realty1 房地产1
type Realty1 struct {
}

//Sell 卖出
func (*Realty1) Sell() {
	fmt.Println("房产1卖出")
}

//Buy 买股票
func (*Realty1) Buy() {
	fmt.Println("房产1买入")
}
