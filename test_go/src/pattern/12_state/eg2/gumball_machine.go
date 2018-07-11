package eg2

import (
	"bytes"
	"fmt"
)

/*
* 新增需求：当曲柄被转动时，有10%的概率掉下的是两颗糖（大赢家）
*/

type GumballMachine struct {
	soldOutState    State //售罄状态
	noQuarterState  State //没有25美分状态
	hasQuarterState State //有25美分状态
	soldState       State //售卖状态
	winnerState     State //大赢家状态

	curState State //当前的状态
	count    int   //剩余的糖果数量
}

func NewGumballMachine(numberGumballs int) *GumballMachine {
	this := new(GumballMachine)
	this.soldOutState = NewSoldOutState(this)
	this.noQuarterState = NewNoQuarterState(this)
	this.hasQuarterState = NewHasQuarterState(this)
	this.soldState = NewSoldState(this)
	this.winnerState = NewWinnerState(this)

	this.curState = this.soldOutState
	this.count = numberGumballs
	if this.count > 0 {
		this.curState = this.noQuarterState
	}
	return this
}

//设置当前的状态
func (this *GumballMachine) SetState(state State) {
	this.curState = state
}

//获取当前的状态
func (this *GumballMachine) GetState() State {
	return this.curState
}

/*----------------------提供获取各个状态的方法------------------------------*/
//获取售罄状态
func (this *GumballMachine) GetSoldOutState() State {
	return this.soldOutState
}

//获取没有25美分的状态
func (this *GumballMachine) GetNoQuarterState() State {
	return this.noQuarterState
}

//获取有25美分的状态
func (this *GumballMachine) GetHasQuarterState() State {
	return this.hasQuarterState
}

//获取售卖状态
func (this *GumballMachine) GetSoldState() State {
	return this.soldState
}

//获取大赢家状态
func (this *GumballMachine) GetWinnerState() State {
	return this.hasQuarterState
}

/*--------------------------操作窗口方法（将动作委托到状态类）-----------------------*/

//插入25美分
func (this *GumballMachine) InsertQuarter() {
	this.curState.InsertQuarter()
}

//退回25美分
func (this *GumballMachine) EjectQuarter() {
	this.curState.EjectQuarter()
}

//转动曲柄
func (this *GumballMachine) TurnCrank() {
	//设计技巧：当状态转换是固定的时候（即，一个状态后必定是另外一个状态），就适合放在Context中。
	// 当状态转换是更动态的时候，通常就会放在状态类中。将状态转换放在状态类中的缺点是：状态类之间产生了依赖。
	this.curState.TurnCrank()
	this.curState.Dispense()
}

//新增加糖
func (this *GumballMachine) Refill(count int) {
	this.count += count
	fmt.Printf("The gumball machine was just refilled; it's new count is: %d\n", this.count)
	this.curState.Refill()
}

/*-------------------------其他方法---------------------------*/

//获取剩余的糖数量
func (this *GumballMachine) GetCount() int {
	return this.count
}

//释放糖果
func (this *GumballMachine) ReleaseBall() {
	fmt.Println("A gumball comes rolling out the slot...") //一颗糖滚出来到一个槽里
	if this.count != 0 {
		this.count -= 1
	}
}

func (this *GumballMachine) String() string {
	buf := bytes.NewBufferString("\nMighty Gumball, Inc.")
	buf.WriteString("\nGo-enabled Standing Gumball Model #2004")
	buf.WriteString("\nInventory: " + fmt.Sprintf("%d", this.count) + " gumball")
	if this.count != 1 {
		buf.WriteString("s")
	}
	buf.WriteString("\n")
	buf.WriteString("Machine is " + this.curState.String() + "\n")
	return buf.String()
}
