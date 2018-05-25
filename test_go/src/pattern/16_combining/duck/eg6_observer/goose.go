package eg6_observer

import "fmt"

// 鹅
type Goose struct {
}

func (this *Goose) Honk() {
	fmt.Println("\nHonk") //鹅的叫声是咯咯
}

func (this *Goose) String() string {
	return "Goose"
}

/*===================================================鹅适配器========================================================*/

// 适配目标：Quackable接口
// 需要实现Quackable接口
type GooseAdapter struct {
	BaseQuackable
	goose *Goose
}

func NewGooseAdapter(goose *Goose) Quackable {
	this := new(GooseAdapter)
	this.BaseQuackable.Init(this)
	this.goose = goose

	return this
}

func (this *GooseAdapter) Quack() {
	this.goose.Honk()
	this.NotifyObserver()
}

func (this *GooseAdapter) String() string {
	return this.goose.String()
}
