package eg2

//产品接口
type Operator interface {
	SetA(int)
	SetB(int)

	Result() int
}

//实现Operator接口实现的基类，封装公用方法（未实现Result方法）
type BaseOperator struct {
	a, b int
}

func (this *BaseOperator) SetA(a int) {
	this.a = a
}

func (this *BaseOperator) SetB(b int) {
	this.b = b
}

/*====================================================================================================================*/

/*
* 1.加法操作
*/
type PlusOperator struct {
	*BaseOperator
}

func NewPlusOperator() *PlusOperator {
	this := new(PlusOperator)
	this.BaseOperator = new(BaseOperator)
	return this
}

func (this *PlusOperator) Result() int {
	return this.a + this.b
}

/*
* 2.减法操作
*/
type MinusOperator struct {
	*BaseOperator
}

func NewMinusOperator() *MinusOperator {
	this := new(MinusOperator)
	this.BaseOperator = new(BaseOperator)
	return this
}

func (this *MinusOperator) Result() int {
	return this.a - this.b
}
