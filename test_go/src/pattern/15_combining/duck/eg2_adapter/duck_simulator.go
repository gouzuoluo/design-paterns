package eg2_adapter

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

	//通过把鹅包装进一个鹅的适配器，使它看起来像一只鸭
	var gooseDuck Quackable = NewGooseAdapter(new(Goose))


	fmt.Println("\nDuck Simulator: With Goose Adapter")

	this.simulate(mallardDuck)
	this.simulate(redheadDuck)
	this.simulate(duckCall)
	this.simulate(rubberDuck)

	this.simulate(gooseDuck)
}

func (this *DuckSimulator) simulate(duck Quackable) {
	duck.Quack()
}
