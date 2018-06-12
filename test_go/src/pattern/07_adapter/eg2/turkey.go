package eg1

import "fmt"

type Turkey interface {
	Gobble() //火鸡只会咯咯叫，不会呱呱叫
}

/*====================================================================================================================*/

type WildTurkey struct {
}

func (this *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}