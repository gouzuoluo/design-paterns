package eg2

import "fmt"

type Coffee struct {
	baseCaffeineBeverage
}

func NewCoffee() *Coffee {
	this := new(Coffee)
	this.CaffeineBeverage = this
	return this
}

func (this *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
