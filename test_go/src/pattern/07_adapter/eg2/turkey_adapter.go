/*
* 类适配器：使用多继承来实现适配功能（天然的双向适配器）
*/
package eg1

//火鸡适配器
type TurkeyAdapter struct {
	Duck
	Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	this := new(TurkeyAdapter)
	this.Turkey = turkey
	return this
}

//重写实现Duck接口的Quack方法
func (this *TurkeyAdapter) Quack() {
	this.Turkey.Gobble()
}