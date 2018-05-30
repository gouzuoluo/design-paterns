package eg2

//工厂接口
type OperatorFactory interface {
	Create() Operator //工厂方法
}

/*====================================================================================================================*/

/*
* 1.PlusOperator 的工厂类
*/
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return NewPlusOperator()
}

/*
* 2.MinusOperator 的工厂类
*/
type MinusOperatorFactory struct {
}

func (MinusOperatorFactory) Create() Operator {
	return NewMinusOperator()
}
