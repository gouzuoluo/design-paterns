package eg6_observer

import "fmt"

//让Quackable接口扩展于QuackObservable接口
type Quackable interface {
	QuackObservable

	Quack()
	String() string
}

//未实现Quack和String方法
type BaseQuackable struct {
	observable *Observable
}

func (this *BaseQuackable) Init(duck Quackable) {
	this.observable = NewObservable(duck)
}

func (this *BaseQuackable) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *BaseQuackable) NotifyObserver() {
	this.observable.NotifyObserver()
}

/*===================================================各种鸭=========================================================*/

/*
* 1.绿头鸭
*/
type MallardDuck struct {
	BaseQuackable
}

func NewMallardDuck() *MallardDuck {
	this := new(MallardDuck)
	this.BaseQuackable.Init(this)

	return this
}

func (this *MallardDuck) Quack() {
	fmt.Println("\nQuack")
	this.NotifyObserver()
}

func (this *MallardDuck) String() string {
	return "Mallard Duck"
}

/*
* 2.红头鸭
*/
type RedHeadDuck struct {
	BaseQuackable
}

func NewRedHeadDuck() *RedHeadDuck {
	this := new(RedHeadDuck)
	this.BaseQuackable.Init(this)

	return this
}

func (this *RedHeadDuck) Quack() {
	fmt.Println("\nQuack")
	this.NotifyObserver()
}

func (this *RedHeadDuck) String() string {
	return "Red Head Duck"
}

/*
* 3.鸭鸣器
*/
type DuckCall struct {
	BaseQuackable
}

func NewDuckCall() *DuckCall {
	this := new(DuckCall)
	this.BaseQuackable.Init(this)

	return this
}

func (this *DuckCall) Quack() {
	fmt.Println("\nKwak")
	this.NotifyObserver()
}

func (this *DuckCall) String() string {
	return "Duck Call"
}

/*
* 4.橡皮鸭
*/
type RubberDuck struct {
	BaseQuackable
}

func NewRubberDuck() *RubberDuck {
	this := new(RubberDuck)
	this.BaseQuackable.Init(this)

	return this
}

func (this *RubberDuck) Quack() {
	fmt.Println("\nSqueak")
	this.NotifyObserver()
}

func (this *RubberDuck) String() string {
	return "Rubber Duck"
}

/*
* 5.诱饵鸭
*/
type DecoyDuck struct {
	BaseQuackable
}

func NewDecoyDuck() *DecoyDuck {
	this := new(DecoyDuck)
	this.BaseQuackable.Init(this)

	return this
}

func (this *DecoyDuck) Quack() {
	fmt.Println("\n<< Silence >>")
	this.NotifyObserver()
}

func (this *DecoyDuck) String() string {
	return "Decoy Duck"
}

/*===================================================呱呱叫统计者=======================================================*/

//QuackCounter是一个装饰者。像一个适配器一样，需要实现目标接口
//目标接口：Quackable接口
var numberOfQuacks int = 0

type QuackCounter struct {
	duck Quackable
}

func NewQuackCounter(duck Quackable) Quackable {
	this := new(QuackCounter)
	this.duck = duck

	return this
}

func (this *QuackCounter) Quack() {
	this.duck.Quack()
	numberOfQuacks++
}

func (this *QuackCounter) RegisterObserver(observer Observer) {
	this.duck.RegisterObserver(observer)
}

func (this *QuackCounter) NotifyObserver() {
	this.duck.NotifyObserver()
}

func (this *QuackCounter) String() string {
	return this.duck.String()
}

func GetQuacks() int {
	return numberOfQuacks
}
