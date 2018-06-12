package eg2

import "fmt"

type Tea struct {
	baseCaffeineBeverage
}

func NewTea() *Tea {
	this := new(Tea)
	this.CaffeineBeverage = this
	return this
}

func (this *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (this *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}