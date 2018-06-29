package eg1

import (
	"fmt"
	"bytes"
)

const (
	SoldOut    = iota //售罄状态
	NoQuarter         //没有25美分状态
	HasQuarter        //有25美分状态
	Sold              //出售状态
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
	this.state = SoldOut
	this.count = count
	if this.count > 0 {
		this.state = NoQuarter
	}
	return this
}

//动作一：投入25美分
func (this *GumballMachine) InsertQuarter() {
	if this.state == HasQuarter {
		fmt.Println("You can't insert another quarter")
	} else if this.state == NoQuarter {
		this.state = HasQuarter
		fmt.Println("You inserted a quarter")
	} else if this.state == SoldOut {
		fmt.Println("You can't insert a quarter, the machine is sold out")
	} else if this.state == Sold {
		fmt.Println("Please wait, we're already giving you a gumball")
	}
}

//动作二：退回25美分
func (this *GumballMachine) EjectQuarter() {
	if this.state == HasQuarter {
		fmt.Println("Quarter returned")
		this.state = NoQuarter
	} else if this.state == NoQuarter {
		fmt.Println("You haven't inserted a quarter")
	} else if this.state == Sold {
		fmt.Println("Sorry, you already turned the crank")
	} else if this.state == SoldOut {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	}
}

//动作三：转动曲柄
func (this *GumballMachine) TurnCrank() {
	if this.state == Sold {
		fmt.Println("Turning twice doesn't get you another gumball!")
	} else if this.state == NoQuarter {
		fmt.Println("You turned but there's no quarter")
	} else if this.state == SoldOut {
		fmt.Println("You turned, but there are no gumballs")
	} else if this.state == HasQuarter {
		fmt.Println("You turned...")
		this.state = Sold
		this.dispense()
	}
}

//动作三：发放糖果
func (this *GumballMachine) dispense() {
	if this.state == Sold {
		fmt.Println("A gumball comes rolling out the slot")
		this.count = this.count - 1
		if this.count == 0 {
			fmt.Println("Oops, out of gumballs!")
			this.state = SoldOut
		} else {
			this.state = NoQuarter
		}
	} else if this.state == NoQuarter {
		fmt.Println("You need to pay first")
	} else if this.state == SoldOut {
		fmt.Println("No gumball dispensed")
	} else if this.state == HasQuarter {
		fmt.Println("No gumball dispensed")
	}
}

//动作四：装入糖果
func (this *GumballMachine) Refill(numGumBalls int) {
	this.count = numGumBalls
	this.state = NoQuarter
}

func (this *GumballMachine) String() string {
	result := bytes.NewBuffer([]byte("\nMighty Gumball, Inc."))
	result.WriteString("\nJava-enabled Standing Gumball Model #2004\n")
	result.WriteString("Inventory: " + fmt.Sprintf("%d", this.count) + " gumball")
	if this.count != 1 {
		result.WriteString("s")
	}
	result.WriteString("\nMachine is ")
	if this.state == SoldOut {
		result.WriteString("sold out")
	} else if this.state == NoQuarter {
		result.WriteString("waiting for quarter")
	} else if this.state == HasQuarter {
		result.WriteString("waiting for turn of crank")
	} else if this.state == Sold {
		result.WriteString("delivering a gumball")
	}
	result.WriteString("\n")
	return result.String()
}
