/*
* 对象适配：使用组合来实现适配功能
*/

package eg1

//客户期望的接口是Duck接口
//我们有的接口是Turkey接口
type TurkeyAdapter struct {
	turkey Turkey //被适配者接口
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	this := new(TurkeyAdapter)
	this.turkey = turkey
	return this
}

//实现Duck接口的Quack方法
func (this *TurkeyAdapter) Quack() {
	this.turkey.Gobble()
}

//实现Duck接口的Fly方法
func (this *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		this.turkey.Fly()
	}
}
