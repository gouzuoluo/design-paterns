package _1_strategy

import "fmt"

type QuackBehavior interface {
	Quack()
}

/*----------------------------------------------------------------------------------------------------------------*/
//1哑巴-不会叫
type MuteQuack struct {
}

func (this *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

//2呱呱叫
type Quack struct {
}

func (this *Quack) Quack() {
	fmt.Println("Quack")
}

//3吱吱叫
type Squeak struct {
}

func (this *Squeak) Quack() {
	fmt.Println("Squeak")
}