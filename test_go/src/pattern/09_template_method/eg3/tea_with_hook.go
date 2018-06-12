package simple_barista

import (
	"fmt"
	"strings"
)

type TeaWithHook struct {
	baseCaffeineBeverageWithHook
}

func NewTeaWithHook() *TeaWithHook {
	this := new(TeaWithHook)
	this.CaffeineBeverageWithHook = this
	return this
}

func (this *TeaWithHook) Brew() {
	fmt.Println("Steeping the tea")
}

func (this *TeaWithHook) AddCondiments() {
	fmt.Println("Adding Lemon")
}

func (this *TeaWithHook) CustomerWantsCondiments() bool {
	answer := this.getUserInput()

	if strings.ToLower(answer) == "y" {
		return true
	} else {
		return false
	}
}

//获取输入
func (this *TeaWithHook) getUserInput() string {
	fmt.Println("Would you like lemon with your tea (y/n)? ")
	answer := ""
	result, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Printf("IO error trying to read your answer,err:%s,result code:%d\n", err, result)
		return "no"
	} else {
		return answer
	}
}
