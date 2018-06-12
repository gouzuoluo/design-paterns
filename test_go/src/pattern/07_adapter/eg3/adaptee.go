/*
* 实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。
*/

package eg3

//被适配的的接口
type Adaptee interface {
	SpecificRequest() string
}

/*==================================================================================================================*/
//是被适配的目标类
type AdapteeImpl struct {
}

func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}
