package eg1

import "fmt"

//鸭子仿真器
type DuckSimulator struct {
}

func NewDuckSimulator()*DuckSimulator  {
	this := new(DuckSimulator)
	return this
}

//仿真
func (this *DuckSimulator) Simulate() {
	var mallardDuck Quackable = new(MallardDuck)
	var redheadDuck Quackable = new(RedHeadDuck)
	var duckCall Quackable = new(DuckCall)
	var rubberDuck Quackable = new(RubberDuck)

	fmt.Println("\nDuck Simulator")

	this.simulate(mallardDuck)
	this.simulate(redheadDuck)
	this.simulate(duckCall)
	this.simulate(rubberDuck)
}

func (this *DuckSimulator) simulate(duck Quackable) {
	duck.Quack()
}
