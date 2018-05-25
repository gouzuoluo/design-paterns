package simple_barista

import "fmt"

type Tea struct {
	baseCaffeineBeverage
}

func NewTea() CaffeineBeverage {
	this := new(Tea)
	this.baseCaffeineBeverage.CaffeineBeverage = this
	return this
}

func (this *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (this *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}