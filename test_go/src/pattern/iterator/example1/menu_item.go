package example1

import "fmt"

//单个菜单项
type MenuItem struct {
	name        string  //菜单项名
	description string  //菜单项描述
	vegetarian  bool    //是否是蔬菜
	price       float64 //价格
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	this := new(MenuItem)
	this.name = name
	this.description = description
	this.vegetarian = vegetarian
	this.price = price
	return this
}

func (this *MenuItem)GetName() string {
	return this.name
}

func (this *MenuItem)GetDescription() string {
	return this.description
}

func (this *MenuItem)GetPrice() float64 {
	return this.price
}

func (this *MenuItem)IsVegetarian() bool {
	return this.vegetarian
}

func (this *MenuItem)String() string {
	return this.name + ", $" + fmt.Sprintf("%f", this.price) + "\n   " + this.description
}