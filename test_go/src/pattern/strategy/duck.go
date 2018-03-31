package strategy

import "fmt"

type Duck interface {
	Swim()
	Display()
	PreformFly()
	PreformQuack()
	SetFlyBehavior(fb FlyBehavior)
	SetQuackBehavior(qb QuackBehavior)
}

type BaseDuck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (this *BaseDuck) PreformFly() {
	this.flyBehavior.Fly()
}

func (this *BaseDuck) PreformQuack() {
	this.quackBehavior.Quack()
}

func (this *BaseDuck) Swim() {
	fmt.Println("All ducks float,even decoys!")
}

func (this *BaseDuck) Display() {
}

func (this *BaseDuck) SetFlyBehavior(fb FlyBehavior) {
	this.flyBehavior = fb
}

func (this *BaseDuck) SetQuackBehavior(qb QuackBehavior) {
	this.quackBehavior = qb
}

/*------------------------------------------------------------------------------*/
//绿头鸭
type MallardDuck struct {
	BaseDuck
}

func NewMallardDuck() *MallardDuck {
	md := new(MallardDuck)
	md.quackBehavior = new(Quack)
	md.flyBehavior = new(FlyWithWings)
	return md
}

func (this *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

//模型鸭
type ModelDuck struct {
	BaseDuck
}

func NewModelDuck() *ModelDuck {
	md := new(ModelDuck)
	md.quackBehavior = new(Quack)
	md.flyBehavior = new(FlyNoWay)
	return md
}

func (this *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}