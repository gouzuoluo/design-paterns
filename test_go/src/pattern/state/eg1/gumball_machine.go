package eg1

import (
	"fmt"
	"bytes"
)

const (
	SOLD_OUT int = iota     //售罄
	NO_QUARTER              //没有25美分
	HAS_QUARTER             //有25美分
	SOLD                    //出售
)

/*
* 糖果机,最初的实现（最简单的思考，如果加一个状态呢?）
*/
type GumballMachine struct {
	state int //这个实例变量跟踪当前状态
	count int //机器内的糖果数量
}

func NewGumballMachine(count int) *GumballMachine {
	this := new(GumballMachine)
	this.state = SOLD_OUT
	this.count = count
	if this.count > 0 {
		this.state = NO_QUARTER
	}
	return this
}

//动作一：投入25美分
func (this *GumballMachine)InsertQuarter() {
	if (this.state == HAS_QUARTER) {
		fmt.Println("You can't insert another quarter")
	} else if (this.state == NO_QUARTER) {
		this.state = HAS_QUARTER
		fmt.Println("You inserted a quarter")
	} else if (this.state == SOLD_OUT) {
		fmt.Println("You can't insert a quarter, the machine is sold out")
	} else if (this.state == SOLD) {
		fmt.Println("Please wait, we're already giving you a gumball")
	}
}

//动作二：退回25美分
func (this *GumballMachine)EjectQuarter() {
	if (this.state == HAS_QUARTER) {
		fmt.Println("Quarter returned")
		this.state = NO_QUARTER
	} else if (this.state == NO_QUARTER) {
		fmt.Println("You haven't inserted a quarter")
	} else if (this.state == SOLD) {
		fmt.Println("Sorry, you already turned the crank")
	} else if (this.state == SOLD_OUT) {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	}
}

//动作三：转动曲柄
func (this *GumballMachine)TurnCrank() {
	if (this.state == SOLD) {
		fmt.Println("Turning twice doesn't get you another gumball!")
	} else if (this.state == NO_QUARTER) {
		fmt.Println("You turned but there's no quarter")
	} else if (this.state == SOLD_OUT) {
		fmt.Println("You turned, but there are no gumballs")
	} else if (this.state == HAS_QUARTER) {
		fmt.Println("You turned...")
		this.state = SOLD
		this.dispense()
	}
}

//动作三：发放糖果(ci)
func (this *GumballMachine)dispense() {
	if (this.state == SOLD) {
		fmt.Println("A gumball comes rolling out the slot")
		this.count = this.count - 1
		if (this.count == 0) {
			fmt.Println("Oops, out of gumballs!")
			this.state = SOLD_OUT
		} else {
			this.state = NO_QUARTER
		}
	} else if (this.state == NO_QUARTER) {
		fmt.Println("You need to pay first")
	} else if (this.state == SOLD_OUT) {
		fmt.Println("No gumball dispensed")
	} else if (this.state == HAS_QUARTER) {
		fmt.Println("No gumball dispensed")
	}
}

//动作四：装入糖果
func (this *GumballMachine)Refill(numGumBalls int) {
	this.count = numGumBalls
	this.state = NO_QUARTER
}

func (this *GumballMachine)String() string {
	result := bytes.NewBuffer([]byte("\nMighty Gumball, Inc."))
	result.WriteString("\nJava-enabled Standing Gumball Model #2004\n")
	result.WriteString("Inventory: " + fmt.Sprintf("%d", this.count) + " gumball")
	if (this.count != 1) {
		result.WriteString("s")
	}
	result.WriteString("\nMachine is ")
	if (this.state == SOLD_OUT) {
		result.WriteString("sold out")
	} else if (this.state == NO_QUARTER) {
		result.WriteString("waiting for quarter")
	} else if (this.state == HAS_QUARTER) {
		result.WriteString("waiting for turn of crank")
	} else if (this.state == SOLD) {
		result.WriteString("delivering a gumball")
	}
	result.WriteString("\n")
	return result.String()
}