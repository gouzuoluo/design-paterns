package strategy

import "fmt"

type QuackBehavior interface {
	Quack()
}

//哑巴-不会叫
type MuteQuack struct {
}

func (this *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

//呱呱叫
type Quack struct {
}

func (this *Quack) Quack() {
	fmt.Println("Quack")
}

//吱吱叫
type Squeak struct {
}

func (this *Squeak) Quack() {
	fmt.Println("Squeak")
}