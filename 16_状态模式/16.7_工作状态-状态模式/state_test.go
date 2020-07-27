package state

func ExampleState() {
	emergencyProjects := NewWork()
	emergencyProjects.Hour = 9
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 10
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 12
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 13
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 14
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 17

	//emergencyProjects.WorkFinished = true
	emergencyProjects.TaskFinished = false

	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 19
	emergencyProjects.WriteProgram()
	emergencyProjects.Hour = 22
	emergencyProjects.WriteProgram()

	// OutPut:
	// 当前时间：9点 上午工作，精神百倍
	// 当前时间：10点 上午工作，精神百倍
	// 当前时间：12点 饿了，午饭；犯困，午休。
	// 当前时间：13点 下午状态还不错，继续努力
	// 当前时间：14点 下午状态还不错，继续努力
	// 当前时间：17点 加班哦，疲累之极
	// 当前时间：19点 加班哦，疲累之极
	// 当前时间：22点 不行了，睡着了。
}
