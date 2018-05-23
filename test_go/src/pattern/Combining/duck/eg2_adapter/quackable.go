package eg2_adapter

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