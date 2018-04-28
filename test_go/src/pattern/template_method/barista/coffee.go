package simple_barista

import "fmt"

type Coffee struct {
	baseCaffeineBeverage
}

func NewCoffee() CaffeineBeverage {
	this := new(Coffee)
	this.baseCaffeineBeverage.CaffeineBeverage = this
	return this
}

func (this *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
