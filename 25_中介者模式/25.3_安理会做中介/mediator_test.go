package mediator

func ExampleMediator() {
	UNSC := &UnitedNationsSecurityCouncil{}

	c1 := NewUSA(UNSC)
	c2 := NewIraq(UNSC)

	UNSC.Colleague1 = c1
	UNSC.Colleague2 = c2

	c1.Declare("不准研制核武器，否则要发动战争！")
	c2.Declare("我们没有核武器，也不怕侵略。")

	// OutPut:
	// 伊拉克获得对方信息：不准研制核武器，否则要发动战争！
	// 美国获得对方信息：我们没有核武器，也不怕侵略。
}
