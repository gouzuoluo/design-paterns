package eg2

/*
* 模板方法：用作一个算法的模板
*/

import "fmt"

//咖啡因饮料接口
type CaffeineBeverage interface {
	BoilWater()     //烧水
	Brew()          //冲泡
	PourInCup()     //倒进杯子
	AddCondiments() //添加调料

	PrepareRecipe() //准备食谱(模板方法)
	//这个模板方法提供了一个框架，可以让其他的咖啡因饮料插进来。新的咖啡因饮料只需要实现自己的方法就可以了。
}

//抽象的咖啡因基类（未实现Brew,AddCondiments方法）
type baseCaffeineBeverage struct {
	CaffeineBeverage
}

//准备食谱
func (this *baseCaffeineBeverage) PrepareRecipe() {
	this.CaffeineBeverage.BoilWater()
	this.CaffeineBeverage.Brew()
	this.CaffeineBeverage.PourInCup()
	this.CaffeineBeverage.AddCondiments()
}

//烧水
func (this *baseCaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

//水倒进杯子
func (this *baseCaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}
