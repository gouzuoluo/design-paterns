package eg2

import (
	"fmt"
	"math/rand"
)

//状态接口
type State interface {
	//插入25美分
	InsertQuarter()
	//退回25美分
	EjectQuarter()
	//转动轴
	TurnCrank()
	//出糖果
	Dispense()
	//重新添加糖果
	Refill()
	//格式化字符串
	String() string
}

/*====================================================================================================================*/

/*
* 1.没有25美分状态
 */
type NoQuarterState struct {
	gumballMachine *GumballMachine //保存一个糖果机器的引用
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	this := new(NoQuarterState)
	this.gumballMachine = gumballMachine

	return this
}

//插入25美分
func (this *NoQuarterState)  InsertQuarter() {
	fmt.Println("You inserted a quarter")
	this.gumballMachine.SetState(this.gumballMachine.GetHasQuarterState())//修改为有25美分状态
}

//退回25美分
func (this *NoQuarterState)  EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

//转动轴
func (this *NoQuarterState)  TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

//出糖果
func (this *NoQuarterState)  Dispense() {
	fmt.Println("You need to pay first")
}

//重新添加糖果
func (this *NoQuarterState)  Refill() {

}

func (this *NoQuarterState)String() string {
	return "waiting for quarter"
}

/*
* 2.有25美分状态
 */
type HasQuarterState struct {
	gumballMachine *GumballMachine //保存一个糖果机器的引用
}

func NewHasQuarterState(gumballMachine *GumballMachine) *HasQuarterState {
	this := new(HasQuarterState)
	this.gumballMachine = gumballMachine

	return this
}

//插入25美分
func (this *HasQuarterState)  InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

//退回25美分
func (this *HasQuarterState)  EjectQuarter() {
	fmt.Println("Quarter returned")
	this.gumballMachine.SetState(this.gumballMachine.GetNoQuarterState())
}

//转动轴
func (this *HasQuarterState)  TurnCrank() {
	fmt.Println("You turned...")
	if v := rand.Intn(10); v == 0 && this.gumballMachine.GetCount() > 1 {
		this.gumballMachine.SetState(this.gumballMachine.GetWinnerState())//修改为大赢家状态
	} else {
		this.gumballMachine.SetState(this.gumballMachine.GetSoldState())//修改为售卖状态
	}
}

//出糖果
func (this *HasQuarterState)  Dispense() {
	fmt.Println("No gumball dispensed")
}

//重新添加糖果
func (this *HasQuarterState)  Refill() {

}

func (this *HasQuarterState)String() string {
	return "waiting for turn of crank"
}

/*
* 3.售卖状态
 */
type SoldState struct {
	gumballMachine *GumballMachine //保存一个糖果机器的引用
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	this := new(SoldState)
	this.gumballMachine = gumballMachine

	return this
}

//插入25美分
func (this *SoldState)  InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

//退回25美分
func (this *SoldState)  EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

//转动轴
func (this *SoldState)  TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

//出糖果
func (this *SoldState)  Dispense() {
	this.gumballMachine.ReleaseBall()
	if this.gumballMachine.GetCount() > 0 {
		this.gumballMachine.SetState(this.gumballMachine.GetNoQuarterState())//如果还有糖果，就设为没有25美分状态
	} else {
		fmt.Println("Oops, out of gumballs!")
		this.gumballMachine.SetState(this.gumballMachine.GetSoldOutState())//如果没有糖果了，就设为售罄状态
	}
}

//重新添加糖果
func (this *SoldState)  Refill() {
}

func (this *SoldState)String() string {
	return "dispensing a gumball"
}


/*
* 4.售罄状态
 */
type SoldOutState struct {
	gumballMachine *GumballMachine //保存一个糖果机器的引用
}

func NewSoldOutState(gumballMachine *GumballMachine) *SoldOutState {
	this := new(SoldOutState)
	this.gumballMachine = gumballMachine

	return this
}

//插入25美分
func (this *SoldOutState)  InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

//退回25美分
func (this *SoldOutState)  EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

//转动轴
func (this *SoldOutState)  TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")

}

//出糖果
func (this *SoldOutState)  Dispense() {
	fmt.Println("No gumball dispensed")
}

//重新添加糖果
func (this *SoldOutState)  Refill() {
	this.gumballMachine.SetState(this.gumballMachine.GetNoQuarterState())//重置状态为没有25美分的状态
}

func (this *SoldOutState)String() string {
	return "sold out"
}


/*
* 5.大赢家状态
 */
type WinnerState struct {
	gumballMachine *GumballMachine //保存一个糖果机器的引用
}

func NewWinnerState(gumballMachine *GumballMachine) *WinnerState {
	this := new(WinnerState)
	this.gumballMachine = gumballMachine

	return this
}

//插入25美分
func (this *WinnerState)  InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

//退回25美分
func (this *WinnerState)  EjectQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

//转动轴
func (this *WinnerState)  TurnCrank() {
	fmt.Println("Turning again doesn't get you another gumball!")
}

//出糖果
func (this *WinnerState)  Dispense() {
	this.gumballMachine.ReleaseBall()//放一颗糖
	if this.gumballMachine.GetCount() == 0 {
		this.gumballMachine.SetState(this.gumballMachine.GetSoldOutState())
	} else {
		this.gumballMachine.ReleaseBall()//再次释放一颗糖
		fmt.Println("YOU'RE A WINNER! You got two gumballs for your quarter")

		if this.gumballMachine.GetCount() > 0 {
			this.gumballMachine.SetState(this.gumballMachine.GetNoQuarterState())
		} else {
			fmt.Println("Oops, out of gumballs!")
			this.gumballMachine.SetState(this.gumballMachine.GetSoldOutState())
		}
	}
}

//重新添加糖果
func (this *WinnerState)  Refill() {
}

func (this *WinnerState)String() string {
	return "despensing two gumballs for your quarter, because YOU'RE A WINNER!"
}