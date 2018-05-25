package _1_strategy

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
	flyBehavior   FlyBehavior //将变化的行为封装（策略模式的精华）
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

func (this *BaseDuck) SetFlyBehavior(fb FlyBehavior) {
	this.flyBehavior = fb
}

func (this *BaseDuck) SetQuackBehavior(qb QuackBehavior) {
	this.quackBehavior = qb
}

//未实现Display方法

/*---------------------------------------------------------------------------------------------------------------*/
//1绿头鸭
type MallardDuck struct {
	BaseDuck
}

func NewMallardDuck() *MallardDuck {
	this := new(MallardDuck)
	this.quackBehavior = new(Quack)
	this.flyBehavior = new(FlyWithWings)
	return this
}

func (_ *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

//2模型鸭
type ModelDuck struct {
	BaseDuck
}

func NewModelDuck() *ModelDuck {
	this := new(ModelDuck)
	this.quackBehavior = new(Quack)
	this.flyBehavior = new(FlyNoWay)
	return this
}

func (_ *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}
