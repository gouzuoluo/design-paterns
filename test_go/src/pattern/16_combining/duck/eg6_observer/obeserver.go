package eg6_observer

import "fmt"

type Observer interface {
	Update(QuackObservable)
}

/*===================================================================================================================*/

/*
* 1.呱呱加学家
*/

type Quackologist struct {
}

func (this *Quackologist) Update(duck QuackObservable) {
	fmt.Printf("Quackologist: %s just quacked.", duck) //简单的打印正在呱呱叫的对象
}

func (this *Quackologist)String()string  {
	return "Quackologist"
}
