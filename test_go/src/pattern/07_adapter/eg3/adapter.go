/*
* 在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，
* 又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口。
*/

package eg3

//转换Adaptee为Target接口的适配器(双向适配器)
type Adapter struct {
	Adaptee
}

//实现Target接口
func (this *Adapter) Request() string {
	return this.SpecificRequest()
}

//Adapter的工厂函数
func NewAdapter(adaptee Adaptee)Target  {
	return &Adapter{
		Adaptee: adaptee,
	}
}