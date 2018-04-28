package simple_barista

import "fmt"

//咖啡因饮料接口
type  CaffeineBeverageWithHook interface {
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

	//新增一个钩子方法
	CustomerWantsCondiments() bool
}


//带钩子的咖啡因饮料
type baseCaffeineBeverageWithHook struct {
	CaffeineBeverageWithHook
}

//准备食谱
func (this *baseCaffeineBeverageWithHook)PrepareRecipe() {
	this.CaffeineBeverageWithHook.BoilWater()
	this.CaffeineBeverageWithHook.Brew()
	this.CaffeineBeverageWithHook.PourInCup()
	if this.CaffeineBeverageWithHook.CustomerWantsCondiments() {
		this.CaffeineBeverageWithHook.AddCondiments()
	}
}

//烧水
func (this *baseCaffeineBeverageWithHook)BoilWater() {
	fmt.Println("Boiling water")
}

//水倒进杯子
func (this *baseCaffeineBeverageWithHook)PourInCup() {
	fmt.Println("Pouring into cup")
}

//给钩子方法一个默认的实现
func (this *baseCaffeineBeverageWithHook)CustomerWantsCondiments() bool {
	return true
}