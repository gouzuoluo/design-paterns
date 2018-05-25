package simple_barista

import (
	"fmt"
	"strings"
)

type CoffeeWithHook struct {
	baseCaffeineBeverageWithHook
}

func NewCoffeeWithHook() CaffeineBeverageWithHook {
	this := new(CoffeeWithHook)
	this.baseCaffeineBeverageWithHook.CaffeineBeverageWithHook = this
	return this
}

func (this *CoffeeWithHook) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *CoffeeWithHook) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}

func (this *CoffeeWithHook)CustomerWantsCondiments() bool {
	answer := this.getUserInput()

	if strings.ToLower(answer) == "y" {
		return true
	} else {
		return false
	}
}

//获取输入
func (this *CoffeeWithHook)getUserInput() string {
	fmt.Println("Would you like lemon with your tea (y/n)? ")
	answer := ""
	result, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Printf("IO error trying to read your answer,err:%s,result code:%d\n",err, result)
		return "no"
	} else {
		return answer
	}
}