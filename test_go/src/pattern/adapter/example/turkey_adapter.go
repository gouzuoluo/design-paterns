package example

//客户期望的接口是Duck接口
//我们有的接口是Turkey接口
type TurkeyAdapter struct {
	turkey Turkey //被适配者接口
}

func NewTurkeyAdapter(turkey Turkey) Duck {
	this := new(TurkeyAdapter)
	this.turkey = turkey
	return this
}

//实现Duck接口的Quack方法
func (this *TurkeyAdapter)Quack() {
	this.turkey.Gobble()
}

//实现Duck接口的Fly方法
func (this *TurkeyAdapter)Fly() {
	for i := 0; i < 5; i++ {
		this.turkey.Fly()
	}
}