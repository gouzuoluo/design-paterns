package _1_strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

/*--------------------------------------------------------------------------------------------------------------*/

//1不会飞行
type FlyNoWay struct {
}

func (this *FlyNoWay)Fly()  {
	fmt.Println("I can't Fly")
}

//2火箭动力飞行
type FlyRocketPowered struct {
}

func (this *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}

//3会飞行
type FlyWithWings struct {
}

func (this *FlyWithWings)Fly()  {
	fmt.Println("I'm flying!!")
}