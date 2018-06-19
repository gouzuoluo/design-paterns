package eg2

import "fmt"

//菜单组件接口
type MenuComponent interface {
	Add(menuComponent MenuComponent)
	Remove(menuComponent MenuComponent)
	GetChild(i int) MenuComponent
	GetName() string
	GetDescription() string
	IsVegetarian() bool
	Print()

	CreateIterator() Iterator //添加一个创建迭代器的方法
}

//基础菜单组件（未实现CreateIterator方法）
type BaseMenuComponent struct {
}

func (this *BaseMenuComponent) Add(menuComponent MenuComponent) {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) Remove(menuComponent MenuComponent) {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) GetChild(i int) MenuComponent {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) GetName() string {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) GetDescription() string {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) IsVegetarian() bool {
	panic("Unsupported Operation")
}

func (this *BaseMenuComponent) Print() {
	panic("Unsupported Operation")
}

/*====================================================================================================================*/

/*
* 1.单个菜单项(继承菜单组件类)
*/
type MenuItem struct {
	BaseMenuComponent

	iterator    Iterator //迭代器
	name        string   //菜单项名
	description string   //菜单项描述
	vegetarian  bool     //是否是蔬菜
	price       float64  //价格
}

func NewMenuItem(name, description string, vegetarian bool, price float64) MenuComponent {
	this := new(MenuItem)
	this.iterator = nil
	this.name = name
	this.description = description
	this.vegetarian = vegetarian
	this.price = price
	return this
}

//重写MenuComponent接口的GetName方法
func (this *MenuItem) GetName() string {
	return this.name
}

//重写MenuComponent接口的GetDescription方法
func (this *MenuItem) GetDescription() string {
	return this.description
}

//重写MenuComponent接口的GetPrice方法
func (this *MenuItem) GetPrice() float64 {
	return this.price
}

//重写MenuComponent接口的IsVegetarian方法
func (this *MenuItem) IsVegetarian() bool {
	return this.vegetarian
}

//重写MenuComponent接口的Print方法
func (this *MenuItem) Print() {
	fmt.Print("  " + this.GetName())
	if this.IsVegetarian() {
		fmt.Print("(v)")
	}
	fmt.Println(", " + fmt.Sprintf("%f", this.GetPrice()))
	fmt.Println("     -- " + this.GetDescription())
}

//实现MenuComponent接口的CreateIterator方法
func (this *MenuItem) CreateIterator() Iterator {
	if this.iterator == nil {
		this.iterator = NewNilIterator()
	}
	return this.iterator
}

/*--------------------------------------------------------------------------------------------------------------------*/

//菜单集合
type MenuGather struct {
	iterator       Iterator
	menuComponents []MenuComponent
}

func NewMenuGather() *MenuGather {
	this := new(MenuGather)
	this.menuComponents = make([]MenuComponent, 0, 10)
	return this
}

func (this *MenuGather) Add(menuComponent MenuComponent) {
	for _, v := range this.menuComponents {
		if v == menuComponent {
			return
		}
	}
	this.menuComponents = append(this.menuComponents, menuComponent)
}

func (this *MenuGather) Remove(menuComponent MenuComponent) {
	for i, v := range this.menuComponents {
		if v == menuComponent {
			this.menuComponents = append(this.menuComponents[:i], this.menuComponents[i+1:]...)
			return
		}
	}
}

func (this *MenuGather) Get(i int) MenuComponent {
	if len(this.menuComponents) > i {
		return this.menuComponents[i]
	} else {
		return nil
	}
}

func (this *MenuGather) Iterator() Iterator {
	if this.iterator == nil {
		this.iterator = NewMenuGatherIterator(this.menuComponents)
	}
	return this.iterator
}

/*
* 2.菜单(继承菜单组件类)——组合
*/
type Menu struct {
	BaseMenuComponent

	iterator    Iterator    //迭代器
	menuGather  *MenuGather //菜单集合
	name        string      //菜单项名
	description string      //菜单项描述
}

func NewMenu(name, description string) MenuComponent {
	this := new(Menu)
	this.iterator = nil
	this.menuGather = NewMenuGather()
	this.name = name
	this.description = description
	return this
}

//重写MenuComponent接口的GetName方法
func (this *Menu) GetName() string {
	return this.name
}

//重写MenuComponent接口的GetDescription方法
func (this *Menu) GetDescription() string {
	return this.description
}

//重写MenuComponent接口的GetChild方法
func (this *Menu) GetChild(i int) MenuComponent {
	return this.menuGather.Get(i)
}

//重写MenuComponent接口的Add方法
func (this *Menu) Add(menuComponent MenuComponent) {
	this.menuGather.Add(menuComponent)
}

//重写MenuComponent接口的Remove方法
func (this *Menu) Remove(menuComponent MenuComponent) {
	this.menuGather.Remove(menuComponent)
}

//重写MenuComponent接口的Print方法
func (this *Menu) Print() {
	fmt.Print("\n" + this.GetName())
	fmt.Println(", " + this.GetDescription())
	fmt.Println("--------------------------")

	//遍历菜单组合
	iterator := this.menuGather.Iterator()
	for iterator.HasNext() {
		menuComponent := iterator.Next()
		menuComponent.Print()
	}
}

//实现MenuComponent接口的CreateIterator方法
func (this *Menu) CreateIterator() Iterator {
	if this.iterator == nil {
		this.iterator = NewCompositeIterator(this.menuGather.Iterator())
	}
	return this.iterator
}
