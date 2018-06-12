package simple_barista

import "fmt"

//咖啡因饮料接口
type CaffeineBeverageWithHook interface {
	BoilWater()     //烧水
	Brew()          //冲泡
	PourInCup()     //倒进杯子
	AddCondiments() //添加调料

	PrepareRecipe() //准备食谱（模板方法）

	CustomerWantsCondiments() bool //新增一个钩子方法
}

//带钩子的咖啡因饮料（未实行部分方法）
type baseCaffeineBeverageWithHook struct {
	CaffeineBeverageWithHook
}

//准备食谱
func (this *baseCaffeineBeverageWithHook) PrepareRecipe() {
	this.CaffeineBeverageWithHook.BoilWater()
	this.CaffeineBeverageWithHook.Brew()
	this.CaffeineBeverageWithHook.PourInCup()
	if this.CaffeineBeverageWithHook.CustomerWantsCondiments() {
		this.CaffeineBeverageWithHook.AddCondiments()
	}
}

//烧水
func (this *baseCaffeineBeverageWithHook) BoilWater() {
	fmt.Println("Boiling water")
}

//水倒进杯子
func (this *baseCaffeineBeverageWithHook) PourInCup() {
	fmt.Println("Pouring into cup")
}

//给钩子方法1个默认的实现
func (this *baseCaffeineBeverageWithHook) CustomerWantsCondiments() bool {
	return true
}
