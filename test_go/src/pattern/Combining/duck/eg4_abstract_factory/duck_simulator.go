package eg4_abstract_factory

import "fmt"

//鸭子仿真器
type DuckSimulator struct {
}

func NewDuckSimulator() *DuckSimulator {
	this := new(DuckSimulator)
	return this
}

//仿真
//调用仿真方法的时候需要传入一个建鸭工厂
func (this *DuckSimulator) Simulate(duckFactory AbstractDuckFactory) {

	var mallardDuck Quackable = duckFactory.CreateMallardDuck()
	var redHeadDuck Quackable = duckFactory.CreateRedHeadDuck()
	var duckCall Quackable = duckFactory.CreateDuckCall()
	var rubberDuck Quackable = duckFactory.CreateRubberDuck()

	//通过把鹅包装进一个鹅的适配器，使它看起来像一只鸭
	var gooseDuck Quackable = NewGooseAdapter(new(Goose))

	fmt.Println("\nDuck Simulator: With Abstract Factory")

	this.simulate(mallardDuck)
	this.simulate(redHeadDuck)
	this.simulate(duckCall)
	this.simulate(rubberDuck)

	this.simulate(gooseDuck)

	fmt.Printf("The ducks quacked %d times\n", GetQuacks())
}

func (this *DuckSimulator) simulate(duck Quackable) {
	duck.Quack()
}
