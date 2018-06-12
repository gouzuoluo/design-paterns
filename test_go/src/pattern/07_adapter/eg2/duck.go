package eg1

import "fmt"

type Duck interface {
	Quack()
}

/*====================================================================================================================*/

type MallardDuck struct {
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack quack")
}