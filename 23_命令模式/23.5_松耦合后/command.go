package command

import "fmt"

//ICommand 命令接口
type ICommand interface {
	//执行命令
	ExcuteCommand()
}

//Command 命令
type Command struct {
	receiver *Barbecuer
}

func NewCommand(receiver *Barbecuer) *Command {
	return &Command{
		receiver: receiver,
	}
}

//BakeMuttonCommand 烤羊肉串命令
type BakeMuttonCommand struct {
	*Command
}

func NewBakeMuttonCommand(c *Barbecuer) *BakeMuttonCommand {
	return &BakeMuttonCommand{
		Command: NewCommand(c),
	}
}

func (b *BakeMuttonCommand) ExcuteCommand() {
	b.Command.receiver.BakeMutton()
}

//BakeChickenWingCommand 烤鸡翅命令
type BakeChickenWingCommand struct {
	*Command
}

func NewBakeChickenWingCommand(c *Barbecuer) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{
		Command: NewCommand(c),
	}
}

func (b *BakeChickenWingCommand) ExcuteCommand() {
	b.Command.receiver.BakeChickenWing()
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
	orders []ICommand
}

//SetOrder 设置订单
func (w *Waiter) SetOrder(command ICommand) {
	w.orders = append(w.orders, command)
}

//CancelOrder 取消订单
func (w *Waiter) CancelOrder(command ICommand) {
	for index, val := range w.orders {
		if val == command {
			w.orders = append(w.orders[:index], w.orders[index+1:]...)
			break
		}
	}
}

//Notify 通知执行
func (w *Waiter) Notify() {
	for _, command := range w.orders {
		command.ExcuteCommand()
	}
}
