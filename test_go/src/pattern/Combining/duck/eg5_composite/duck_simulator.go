package eg5_composite

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

	fmt.Println("\nDuck Simulator: With Composite - Flocks")

	//先创建一个*Flock主群，然后把所有的Quackable都塞给它
	var flockOfDucks *Flock = NewFlock()
	flockOfDucks.Add(mallardDuck)
	flockOfDucks.Add(redHeadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	//再创建一个新的绿头鸭群
	var flockOfMallards *Flock = NewFlock()
	var mallardDuck1 Quackable = duckFactory.CreateMallardDuck()
	var mallardDuck2 Quackable = duckFactory.CreateMallardDuck()
	var mallardDuck3 Quackable = duckFactory.CreateMallardDuck()
	var mallardDuck4 Quackable = duckFactory.CreateMallardDuck()
	flockOfMallards.Add(mallardDuck1)
	flockOfMallards.Add(mallardDuck2)
	flockOfMallards.Add(mallardDuck3)
	flockOfMallards.Add(mallardDuck4)

	//将绿头鸭群，加入主群
	flockOfDucks.Add(flockOfMallards)

	//测试主群
	fmt.Println("\nDuck Simulator: Whole Flock Simulation")
	this.simulate(flockOfDucks)

	//测试绿头鸭群
	fmt.Println("\nDuck Simulator: Mallard Flock Simulation")
	this.simulate(flockOfMallards)

	fmt.Printf("\nThe ducks quacked %d times\n", GetQuacks())
}

func (this *DuckSimulator) simulate(duck Quackable) {
	duck.Quack()
}
