package eg1

import "fmt"

type Tea struct {

}

func (this *Tea)PrepareRecipe() {
	this.BoilWater()
	this.SteepTeaBag()
	this.PourInCup()
	this.AddLemon()
}

func (this *Tea) BoilWater() {
	fmt.Println("Boiling water")
}

func (this *Tea) SteepTeaBag() {
	fmt.Println("Steeping the tea")
}

func (this *Tea) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (this *Tea) AddLemon() {
	fmt.Println("Adding Lemon")
}