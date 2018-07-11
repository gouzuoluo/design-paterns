package eg3

import "fmt"

//星期接口（状态接口）
type Week interface {
	Today()
	Next(*DayContext)
}

/*===================================================================================================================*/

//1.星期日
type Sunday struct {
}

func (*Sunday) Today() {
	fmt.Sprintf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = new(Monday)
}

//2.星期一
type Monday struct {
}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = new(Tuesday)
}

//3.星期二
type Tuesday struct {
}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = new(Wednesday)
}

//4.星期三
type Wednesday struct {
}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = new(Thursday)
}

//5.星期四
type Thursday struct {
}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = new(Friday)
}

//6.星期五
type Friday struct {
}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = new(Saturday)
}

//7.星期六
type Saturday struct {
}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = new(Sunday)
}
