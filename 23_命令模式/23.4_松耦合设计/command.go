package command

import "fmt"

//ICommand 命令接口
type ICommand interface {
	//执行命令
	ExcuteCommand()
}

//BakeMuttonCommand 烤羊肉串命令
type BakeMuttonCommand struct {
	receiver *Barbecuer
}

func NewBakeMuttonCommand(c *Barbecuer) *BakeMuttonCommand {
	return &BakeMuttonCommand{
		receiver: c,
	}
}

func (b *BakeMuttonCommand) ExcuteCommand() {
	b.receiver.BakeMutton()
}

//BakeChickenWingCommand 烤鸡翅命令
type BakeChickenWingCommand struct {
	receiver *Barbecuer
}

func NewBakeChickenWingCommand(c *Barbecuer) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{
		receiver: c,
	}
}

func (b *BakeChickenWingCommand) ExcuteCommand() {
	b.receiver.BakeChickenWing()
}

//Barbecuer 烤肉串者
type Barbecuer struct {
}

func (*Barbecuer) BakeMutton() {
	fmt.Println("烤羊肉串!")
}

func (*Barbecuer) BakeChickenWing() {
	fmt.Println("烤鸡翅!")
}

//Waiter 服务员
type Waiter struct {
	command ICommand
}

//SetOrder 设置订单
func (w *Waiter) SetOrder(command ICommand) {
	w.command = command
}

//Notify 通知执行
func (w *Waiter) Notify() {
	w.command.ExcuteCommand()
}
