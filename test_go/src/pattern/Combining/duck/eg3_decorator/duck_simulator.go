package eg3_decorator

import "fmt"

//鸭子仿真器
type DuckSimulator struct {
}

func NewDuckSimulator() *DuckSimulator {
	this := new(DuckSimulator)
	return this
}

//仿真
func (this *DuckSimulator) Simulate() {
	var mallardDuck Quackable = NewQuackCounter(new(MallardDuck))
	var redheadDuck Quackable = NewQuackCounter(new(RedHeadDuck))
	var duckCall Quackable = NewQuackCounter(new(DuckCall))
	var rubberDuck Quackable = NewQuackCounter(new(RubberDuck))

	//通过把鹅包装进一个鹅的适配器，使它看起来像一只鸭
	var gooseDuck Quackable = NewGooseAdapter(new(Goose))

	fmt.Println("\nDuck Simulator: With Decorator")

	this.simulate(mallardDuck)
	this.simulate(redheadDuck)
	this.simulate(duckCall)
	this.simulate(rubberDuck)

	this.simulate(gooseDuck)

	fmt.Printf("The ducks quacked %d times\n", GetQuacks())
}

func (this *DuckSimulator) simulate(duck Quackable) {
	duck.Quack()
}
