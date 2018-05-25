package eg3_decorator

import "fmt"

// 鹅
type Goose struct {
}

func (this *Goose) Honk() {
	fmt.Println("Honk") //鹅的叫声是咯咯
}

/*===================================================鹅适配器========================================================*/

// 适配目标：Quackable接口
// 需要实现Quackable接口
type GooseAdapter struct {
	goose *Goose
}

func NewGooseAdapter(goose *Goose) Quackable {
	this := new(GooseAdapter)
	this.goose = goose

	return this
}

func (this *GooseAdapter) Quack() {
	this.goose.Honk()
}
