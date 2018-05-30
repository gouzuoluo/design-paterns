package eg1

/*
*  披萨（抽象工厂派生出的具体工厂的客户就是各种类型的pizza）
*/

import (
	"fmt"
	"bytes"
)

type Pizza interface {
	//在准备阶段，每个具体的披萨，使用的原料种类及分量不同
	Prepare()

	Bake()
	Cut()
	Box()
	SetName(string)
	GetName() string
}

//实现Pizza接口的大部分方法（未实现Prepare方法，由具体的披萨子类实现）
type basePizza struct {
	name      string    //名字
	dough     Dough     //生面团
	sauce     Sauce     //酱油
	veggies   []Veggies //蔬菜
	cheese    Cheese    //奶酪
	pepperoni Pepperoni //意大利香肠
	clam      Clams     //蛤蜊

	ingredientFactory PizzaIngredientFactory //披萨原料工厂（使用抽象工厂，将从实际工厂解耦）
}

func (this *basePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (this *basePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (this *basePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (this *basePizza) SetName(n string) {
	this.name = n
}

func (this *basePizza) GetName() string {
	return this.name
}

//增加一个String方法
func (this *basePizza) String() string {
	buf := bytes.NewBufferString("---- " + this.name + " ----\n")
	if this.dough != nil {
		buf.WriteString(this.dough.String())
		buf.WriteString("\n")
	}
	if this.sauce != nil {
		buf.WriteString(this.sauce.String())
		buf.WriteString("\n")
	}
	if this.cheese != nil {
		buf.WriteString(this.cheese.String())
		buf.WriteString("\n")
	}
	if this.veggies != nil {
		for _, v := range this.veggies {
			buf.WriteString(v.String())
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	if this.clam != nil {
		buf.WriteString(this.clam.String())
		buf.WriteString("\n")
	}
	if this.pepperoni != nil {
		buf.WriteString(this.pepperoni.String())
		buf.WriteString("\n")
	}
	return buf.String()
}

/*------------------------------------------------------------------------------------------------------------------*/

/*
* 1奶油披萨
 */
type CheesePizza struct {
	basePizza
}

func NewCheesePizza(ingredientFactory PizzaIngredientFactory) Pizza {
	this := new(CheesePizza)
	this.ingredientFactory = ingredientFactory
	return this
}

func (this *CheesePizza) Prepare() {
	fmt.Println("Preparing " + this.name)
	this.dough = this.ingredientFactory.CreateDough()
	this.sauce = this.ingredientFactory.CreateSauce()
	this.cheese = this.ingredientFactory.CreateCheese()
}

/*
* 2.蛤蜊披萨
 */
type ClamPizza struct {
	basePizza
}

func NewClamPizza(ingredientFactory PizzaIngredientFactory) Pizza {
	this := new(ClamPizza)
	this.ingredientFactory = ingredientFactory
	return this
}

func (this *ClamPizza) Prepare() {
	fmt.Println("Preparing " + this.name)
	this.dough = this.ingredientFactory.CreateDough()
	this.sauce = this.ingredientFactory.CreateSauce()
	this.cheese = this.ingredientFactory.CreateCheese()
	this.clam = this.ingredientFactory.CreateClam()
}

/*
* 3.意大利香肠披萨
 */
type PepperoniPizza struct {
	basePizza
}

func NewPepperoniPizza(ingredientFactory PizzaIngredientFactory) Pizza {
	this := new(PepperoniPizza)
	this.ingredientFactory = ingredientFactory
	return this
}

func (this *PepperoniPizza) Prepare() {
	fmt.Println("Preparing " + this.name)
	this.dough = this.ingredientFactory.CreateDough()
	this.sauce = this.ingredientFactory.CreateSauce()
	this.cheese = this.ingredientFactory.CreateCheese()
	this.pepperoni = this.ingredientFactory.CreatePepperoni()
}

/*
* 4.蔬菜披萨
 */
type VeggiesPizza struct {
	basePizza
}

func NewVeggiePizza(ingredientFactory PizzaIngredientFactory) Pizza {
	this := new(VeggiesPizza)
	this.ingredientFactory = ingredientFactory
	return this
}

func (this *VeggiesPizza) Prepare() {
	fmt.Println("Preparing " + this.name)
	this.dough = this.ingredientFactory.CreateDough()
	this.sauce = this.ingredientFactory.CreateSauce()
	this.cheese = this.ingredientFactory.CreateCheese()
	this.veggies = this.ingredientFactory.CreateVeggies()
}
