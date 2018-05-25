package simple_barista

import "fmt"

//咖啡因饮料接口
type  CaffeineBeverage interface {
	//烧水
	BoilWater()
	//冲泡
	Brew()
	//倒进杯子
	PourInCup()
	//添加调料
	AddCondiments()

	//准备食谱
	PrepareRecipe()
}

//抽象的咖啡因基类（实现部分方法，未实现的方法就是抽象方法）
type baseCaffeineBeverage struct {
	CaffeineBeverage
}

//准备食谱
func (this *baseCaffeineBeverage)PrepareRecipe() {
	this.CaffeineBeverage.BoilWater()
	this.CaffeineBeverage.Brew()
	this.CaffeineBeverage.PourInCup()
	this.CaffeineBeverage.AddCondiments()
}

//烧水
func (this *baseCaffeineBeverage)BoilWater() {
	fmt.Println("Boiling water")
}

//水倒进杯子
func (this *baseCaffeineBeverage)PourInCup() {
	fmt.Println("Pouring into cup")
}