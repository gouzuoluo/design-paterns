package example

import "fmt"

type Turkey interface {
	Gobble()//火鸡只会咯咯叫，不会呱呱叫
	Fly()//火鸡会飞，虽然飞不远
}

type WildTurkey struct {

}

func (this *WildTurkey)Gobble() {
	fmt.Println("Gobble gobble")
}

func (this *WildTurkey)Fly() {
	fmt.Println("I`m flying a short distance")
}