package templatemethod

import (
	"fmt"
)

//TestPaperA 学生甲抄的试卷
type TestPaperA struct {
}

//TestQuestion1 试题1
func (*TestPaperA) TestQuestion1() {
	fmt.Println(" 杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] a.球磨铸铁 b.马口铁 c.高速合金钢 d.碳素纤维")
	fmt.Println("答案：b")
}

//TestQuestion2 试题2
func (*TestPaperA) TestQuestion2() {
	fmt.Println(" 杨过、程英、陆无双铲除了情花，造成[ ] a.使这种植物不再害人 b.使一种珍稀物种灭绝 c.破坏了那个生物圈的生态平衡 d.造成该地区沙漠化")
	fmt.Println("答案：a")
}

//TestQuestion3 试题3
func (*TestPaperA) TestQuestion3() {
	fmt.Println(" 蓝凤凰的致使华山师徒、桃谷六仙呕吐不止,如果你是大夫,会给他们开什么药[ ] a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量的生牛奶 e.以上全不对")
	fmt.Println("答案：c")
}

//TestPaperB 学生甲抄的试卷
type TestPaperB struct {
}

//TestQuestion1 试题1
func (*TestPaperB) TestQuestion1() {
	fmt.Println(" 杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] a.球磨铸铁 b.马口铁 c.高速合金钢 d.碳素纤维")
	fmt.Println("答案：d")
}

//TestQuestion2 试题2
func (*TestPaperB) TestQuestion2() {
	fmt.Println(" 杨过、程英、陆无双铲除了情花，造成[ ] a.使这种植物不再害人 b.使一种珍稀物种灭绝 c.破坏了那个生物圈的生态平衡 d.造成该地区沙漠化")
	fmt.Println("答案：b")
}

//TestQuestion3 试题3
func (*TestPaperB) TestQuestion3() {
	fmt.Println(" 蓝凤凰的致使华山师徒、桃谷六仙呕吐不止,如果你是大夫,会给他们开什么药[ ] a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量的生牛奶 e.以上全不对")
	fmt.Println("答案：a")
}
