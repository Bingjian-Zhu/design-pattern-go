package facade

import "fmt"

//Fund 基金
type Fund struct {
	gu1 *Stock1
	gu2 *Stock2
	gu3 *Stock3
	nd1 *NationalDebt1
	rt1 *Realty1
}

//NewFund Fund构造函数
func NewFund() *Fund {
	return &Fund{
		gu1: &Stock1{},
		gu2: &Stock2{},
		gu3: &Stock3{},
		nd1: &NationalDebt1{},
		rt1: &Realty1{},
	}
}

//BuyFund 买基金
func (f *Fund) BuyFund() {
	f.gu1.Buy()
	f.gu2.Buy()
	f.gu3.Buy()
	f.nd1.Buy()
	f.rt1.Buy()
}

//SellFund 卖基金
func (f *Fund) SellFund() {
	f.gu1.Sell()
	f.gu2.Sell()
	f.gu3.Sell()
	f.nd1.Sell()
	f.rt1.Sell()
}

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
