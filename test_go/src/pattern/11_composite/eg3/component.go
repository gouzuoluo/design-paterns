package eg3

import "fmt"

//组件接口
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

type BaseComponent struct {
	parent Component
	name   string
}

func (this *BaseComponent) Parent() Component {
	return &this.parent
}

func (this *BaseComponent) SetParent(parent Component) {
	this.parent = parent
}

func (this *BaseComponent) Name() string {
	return this.name
}

func (this *BaseComponent) SetName(name string) {
	this.name = name
}

func (this *BaseComponent) AddChild(Component) {

}

func (this *BaseComponent) Print(string) {

}

/*====================================================================================================================*/

//1.叶子组件
type Leaf struct {
	BaseComponent
}

func NewLeaf() *Leaf {
	this := new(Leaf)
	return this
}

func (this *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, this.Name())
}

//2.组合组件
type Composite struct {
	BaseComponent
	children []Component
}

func NewComposite() *Composite {
	this := new(Composite)
	this.children = make([]Component, 0)
	return this
}

func (this *Composite) AddChild(child Component) {
	child.SetParent(this)
	this.children = append(this.children, child)
}

func (this *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, this.Name())
	pre += " "
	for _, comp := range this.children {
		comp.Print(pre)
	}
}
