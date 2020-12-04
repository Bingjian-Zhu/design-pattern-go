package adapter

func ExampleAdapter() {
	b := new(Forwards)
	b.Name = "巴蒂尔"
	b.Attack()

	m := new(Guards)
	m.Name = "麦克格雷迪"
	m.Attack()

	ym := NewTranslator("姚明")
	ym.Attack()
	ym.Defense()

	// OutPut:
	// 前锋 巴蒂尔 进攻
	// 后卫 麦克格雷迪 进攻
	// 外籍中锋 姚明 进攻
	// 外籍中锋 姚明 防守
}
