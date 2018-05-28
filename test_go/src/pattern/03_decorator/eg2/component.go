package eg2

//组件接口
type Component interface {
	Calc() int //计算
}

/*===================================================具体组件========================================================*/

type ConcreteComponent struct {
}

func NewConcreteComponent() *ConcreteComponent {
	this := new(ConcreteComponent)
	return this
}

func (*ConcreteComponent) Calc() int {
	return 0
}
