package eg6_observer

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
	return NewMallardDuck()
}

//创建红头鸭
func (this *DuckFactory) CreateRedHeadDuck() Quackable {
	return NewRedHeadDuck()
}

//创建鸭鸣器
func (this *DuckFactory) CreateDuckCall() Quackable {
	return NewDuckCall()
}

//创建橡皮鸭
func (this *DuckFactory) CreateRubberDuck() Quackable {
	return NewRubberDuck()
}

//创建诱饵鸭
func (this *DuckFactory) CreateDecoyDuck() Quackable {
	return NewDecoyDuck()
}


/*
* 2.创建带有装饰者鸭的工厂
*/
type CountingDuckFactory struct {
}

//创建绿头鸭
func (this *CountingDuckFactory) CreateMallardDuck() Quackable {
	return NewQuackCounter(NewMallardDuck())
}

//创建红头鸭
func (this *CountingDuckFactory) CreateRedHeadDuck() Quackable {
	return NewQuackCounter(NewRedHeadDuck())
}

//创建鸭鸣器
func (this *CountingDuckFactory) CreateDuckCall() Quackable {
	return NewQuackCounter(NewDuckCall())
}

//创建橡皮鸭
func (this *CountingDuckFactory) CreateRubberDuck() Quackable {
	return NewQuackCounter(NewRubberDuck())
}

//创建诱饵鸭
func (this *CountingDuckFactory) CreateDecoyDuck() Quackable {
	return NewQuackCounter(NewDecoyDuck())
}
