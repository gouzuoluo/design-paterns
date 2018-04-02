package decorator

//实现Beverage部分方法
type Condiment struct {
	beverage Beverage //存放被装饰的对象的引用
}

func (this *Condiment) GetSize() Size {
	return this.beverage.GetSize()
}

func (this *Condiment) SetSize(size Size) {
	this.beverage.SetSize(size)
}

//未实现GetDescription方法

//未实现Cost方法

/*---------------------------------------------------调味品---------------------------------------------------------*/
//调味品1
type Milk struct {
	Condiment
}

func NewCondiment(beverage Beverage) Beverage {
	this := new(Condiment)
	this.beverage = beverage
	return this
}

func (this *Condiment) GetDescription() string {
	return this.beverage.GetDescription() + ",Milk"
}

func (this *Condiment) Cost() float64 {
	return this.beverage.Cost() + .10
}

//调味品2
type Mocha struct {
	Condiment
}

func NewMocha(beverage Beverage) Beverage {
	this := new(Mocha)
	this.beverage = beverage
	return this
}

func (this *Mocha) GetDescription() string {
	return this.beverage.GetDescription() + ",Mocha"
}

func (this *Mocha) Cost() float64 {
	return this.beverage.Cost() + .20
}

//调味品3
type Soy struct {
	Condiment
}

func NewSoy(beverage Beverage) Beverage {
	this := new(Soy)
	this.beverage = beverage
	return this
}

func (this *Soy) GetDescription() string {
	return this.beverage.GetDescription() + ",Soy"
}

func (this *Soy) Cost() float64 {
	cost := this.beverage.Cost()
	if this.beverage.GetSize() == TALL {
		cost += .10
	} else if this.beverage.GetSize() == GRANDE {
		cost += .15
	} else if this.beverage.GetSize() == VENTI {
		cost += .20
	}
	return cost
}

//调味品4
type Whip struct {
	Condiment
}

func NewWhip(beverage Beverage) Beverage {
	this := new(Whip)
	this.beverage = beverage
	return this
}

func (this *Whip) GetDescription() string {
	return this.beverage.GetDescription() + ",Wip"
}

func (this *Whip) Cost() float64 {
	return this.beverage.Cost() + .10
}
