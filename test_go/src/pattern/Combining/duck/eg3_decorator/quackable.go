package eg3_decorator

import "fmt"

type Quackable interface {
	Quack()
}

/*===================================================各种鸭=========================================================*/

/*
* 1.绿头鸭
*/
type MallardDuck struct {
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack")
}

/*
* 2.红头鸭
*/
type RedHeadDuck struct {
}

func (this *RedHeadDuck) Quack() {
	fmt.Println("Quack")
}

/*
* 3.鸭鸣器
*/
type DuckCall struct {
}

func (this *DuckCall) Quack() {
	fmt.Println("Kwak")
}

/*
* 4.橡皮鸭
*/
type RubberDuck struct {
}

func (this *RubberDuck) Quack() {
	fmt.Println("Squeak")
}

/*
* 5.诱饵鸭
*/
type DecoyDuck struct {
}

func (this *DecoyDuck) Quack() {
	fmt.Println("<< Silence >>")
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

func GetQuacks() int {
	return numberOfQuacks
}
