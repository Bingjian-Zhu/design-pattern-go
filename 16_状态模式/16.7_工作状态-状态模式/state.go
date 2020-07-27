package state

import "fmt"

//State 抽象状态
type State interface {
	WriteProgram(w *Work)
}

//ForenoonState 上午工作状态
type ForenoonState struct {
}

//WriteProgram 工作状态
func (*ForenoonState) WriteProgram(w *Work) {
	if w.Hour < 12 {
		fmt.Printf("当前时间：%g点 上午工作，精神百倍\n", w.Hour)
	} else {
		w.SetState(&NoonState{})
		w.WriteProgram()
	}
}

//NoonState 中午工作状态
type NoonState struct {
}

//WriteProgram 工作状态
func (*NoonState) WriteProgram(w *Work) {
	if w.Hour < 13 {
		fmt.Printf("当前时间：%g点 饿了，午饭；犯困，午休。\n", w.Hour)
	} else {
		w.SetState(&AfternoonState{})
		w.WriteProgram()
	}
}

//AfternoonState 下午工作状态
type AfternoonState struct {
}

//WriteProgram 工作状态
func (*AfternoonState) WriteProgram(w *Work) {
	if w.Hour < 17 {
		fmt.Printf("当前时间：%g点 下午状态还不错，继续努力\n", w.Hour)
	} else {
		w.SetState(&EveningState{})
		w.WriteProgram()
	}
}

//EveningState 晚间工作状态
type EveningState struct {
}

//WriteProgram 工作状态
func (*EveningState) WriteProgram(w *Work) {
	if w.TaskFinished {
		w.SetState(&RestState{})
		w.WriteProgram()
	} else {
		if w.Hour < 21 {
			fmt.Printf("当前时间：%g点 加班哦，疲累之极\n", w.Hour)
		} else {
			w.SetState(&SleepingState{})
			w.WriteProgram()
		}
	}
}

//SleepingState 睡眠状态
type SleepingState struct {
}

//WriteProgram 工作状态
func (*SleepingState) WriteProgram(w *Work) {
	fmt.Printf("当前时间：%g点 不行了，睡着了。\n", w.Hour)
}

//RestState 下班休息状态
type RestState struct {
}

//WriteProgram 工作状态
func (*RestState) WriteProgram(w *Work) {
	fmt.Printf("当前时间：%g点 下班回家了\n", w.Hour)
}

//Work 工作
type Work struct {
	current      State
	Hour         float64
	TaskFinished bool
}

// NewWork 构造函数
func NewWork() *Work {
	return &Work{
		current: &ForenoonState{},
	}
}

//SetState 设置状态
func (w *Work) SetState(s State) {
	w.current = s
}

//WriteProgram 工作状态
func (w *Work) WriteProgram() {
	w.current.WriteProgram(w)
}
