package eg4_abstract_factory

type AbstractDuckFactory interface {
	//创建绿头鸭
	CreateMallardDuck() Quackable

	//创建红头鸭
	CreateRedHeadDuck() Quackable

	//创建鸭鸣器
	CreateDuckCall() Quackable

	//创建橡皮鸭
	CreateRubberDuck() Quackable

	//创建诱饵鸭
	CreateDecoyDuck() Quackable
}

/*=======================================================创建鸭工厂===================================================*/

/*
* 1.创建没有装饰者鸭的工厂
*/
type DuckFactory struct {
}

//创建绿头鸭
func (this *DuckFactory) CreateMallardDuck() Quackable {
	return new(MallardDuck)
}

//创建红头鸭
func (this *DuckFactory) CreateRedHeadDuck() Quackable {
	return new(RedHeadDuck)
}

//创建鸭鸣器
func (this *DuckFactory) CreateDuckCall() Quackable {
	return new(DuckCall)
}

//创建橡皮鸭
func (this *DuckFactory) CreateRubberDuck() Quackable {
	return new(RubberDuck)
}

//创建诱饵鸭
func (this *DuckFactory) CreateDecoyDuck() Quackable {
	return new(DecoyDuck)
}


/*
* 2.创建带有装饰者鸭的工厂
*/
type CountingDuckFactory struct {
}

//创建绿头鸭
func (this *CountingDuckFactory) CreateMallardDuck() Quackable {
	return NewQuackCounter(new(MallardDuck))
}

//创建红头鸭
func (this *CountingDuckFactory) CreateRedHeadDuck() Quackable {
	return NewQuackCounter(new(RedHeadDuck))
}

//创建鸭鸣器
func (this *CountingDuckFactory) CreateDuckCall() Quackable {
	return NewQuackCounter(new(DuckCall))
}

//创建橡皮鸭
func (this *CountingDuckFactory) CreateRubberDuck() Quackable {
	return NewQuackCounter(new(RubberDuck))
}

//创建诱饵鸭
func (this *CountingDuckFactory) CreateDecoyDuck() Quackable {
	return NewQuackCounter(new(DecoyDuck))
}
