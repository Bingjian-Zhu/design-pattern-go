package templatemethodc

import "fmt"

//IAnswer 答案接口
type IAnswer interface {
	Answer1() string
	Answer2() string
	Answer3() string
}

//TestPaper 金庸小说考题试卷
type TestPaper struct {
	IAnswer
}

//NewTestPaper TestPaper构造函数
func NewTestPaper(answer IAnswer) *TestPaper {
	return &TestPaper{
		IAnswer: answer,
	}
}

//TestQuestion1 试题1
func (t *TestPaper) TestQuestion1() {
	fmt.Println(" 杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] a.球磨铸铁 b.马口铁 c.高速合金钢 d.碳素纤维")
	fmt.Printf("答案：%s\n", t.IAnswer.Answer1())
}

//TestQuestion2 试题2
func (t *TestPaper) TestQuestion2() {
	fmt.Println(" 杨过、程英、陆无双铲除了情花，造成[ ] a.使这种植物不再害人 b.使一种珍稀物种灭绝 c.破坏了那个生物圈的生态平衡 d.造成该地区沙漠化")
	fmt.Printf("答案：%s\n", t.IAnswer.Answer2())
}

//TestQuestion3 试题3
func (t *TestPaper) TestQuestion3() {
	fmt.Println(" 蓝凤凰的致使华山师徒、桃谷六仙呕吐不止,如果你是大夫,会给他们开什么药[ ] a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量的生牛奶 e.以上全不对")
	fmt.Printf("答案：%s\n", t.IAnswer.Answer3())
}

//TestPaperA 试卷A
type TestPaperA struct {
}

//Answer1 答案1
func (*TestPaperA) Answer1() string {
	return "b"
}

//Answer2 答案2
func (*TestPaperA) Answer2() string {
	return "c"
}

//Answer3 答案3
func (*TestPaperA) Answer3() string {
	return "a"
}

//TestPaperB 试卷B
type TestPaperB struct {
}

//Answer1 答案1
func (*TestPaperB) Answer1() string {
	return "c"
}

//Answer2 答案2
func (*TestPaperB) Answer2() string {
	return "a"
}

//Answer3 答案3
func (*TestPaperB) Answer3() string {
	return "a"
}
