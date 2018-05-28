package eg2

//装饰者1
type MulDecorator struct {
	component Component
	num       int
}

func NewMulDecorator(component Component, num int) *MulDecorator {
	this := new(MulDecorator)
	this.component = component
	this.num = num
	return this
}

func (this *MulDecorator) Calc() int {
	return this.component.Calc() * this.num
}

//装饰者2
type AddDecorator struct {
	component Component
	num       int
}

func NewAddDecorator(component Component, num int) *AddDecorator {
	this := new(AddDecorator)
	this.component = component
	this.num = num
	return this
}

func (this *AddDecorator) Calc() int {
	return this.component.Calc() + this.num
}
