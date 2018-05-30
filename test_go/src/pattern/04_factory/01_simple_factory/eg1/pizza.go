package eg1

import (
	"fmt"
	"bytes"
)

//产品接口
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

//实现Pizza所有接口
type BasePizza struct {
	name     string   //名字
	dough    string   //生面团
	sauce    string   //酱油
	toppings []string //配料
}

func (this *BasePizza) Prepare() {
	fmt.Println("Preparing ", this.name)
}

func (this *BasePizza) Bake() {
	fmt.Println("Baking ", this.name)
}

func (this *BasePizza) Cut() {
	fmt.Println("Cuting ", this.name)
}

func (this *BasePizza) Box() {
	fmt.Println("Boxing ", this.name)
}

func (this *BasePizza) GetName() string {
	return this.name
}

//实现一个String方法，用于打印
func (this *BasePizza) String() string {
	buf := bytes.NewBufferString("---- " + this.name + " ----\n")
	buf.Write([]byte(this.dough + "\n"))
	buf.Write([]byte(this.sauce + "\n"))
	for _, v := range this.toppings {
		buf.Write([]byte(v + "\n"))
	}
	return buf.String()
}

/*====================================================================================================================*/

/*
* 1奶酪披萨
 */
type CheesePizza struct {
	BasePizza
}

func NewCheesePizza() Pizza {
	this := new(CheesePizza)
	this.name = "Cheese Pizza"
	this.dough = "Regular Crust"
	this.sauce = "Marinara Pizza Sauce"
	this.toppings = append(this.toppings, "Fresh Mozzarella")
	this.toppings = append(this.toppings, "Parmesan")
	return this
}

/*
* 2.蛤蜊披萨
 */
type ClamPizza struct {
	BasePizza
}

func NewClamPizza() Pizza {
	this := new(ClamPizza)
	this.name = "Clam Pizza"
	this.dough = "Thin crust"
	this.sauce = "White garlic sauce"
	this.toppings = append(this.toppings, "Clams")
	this.toppings = append(this.toppings, "Grated parmesan cheese")
	return this
}

/*
* 3.意大利辣香肠披萨
*/
type PepperoniPizza struct {
	BasePizza
}

func NewPepperoniPizza() Pizza {
	this := new(PepperoniPizza)
	this.name = "Pepperoni Pizza"
	this.dough = "Crust"
	this.sauce = "Marinara sauce"
	this.toppings = append(this.toppings, "Sliced Pepperoni")
	this.toppings = append(this.toppings, "Sliced Onion")
	this.toppings = append(this.toppings, "Grated parmesan cheese")
	return this
}

/*
* 4.蔬菜披萨
*/
type VeggiePizza struct {
	BasePizza
}

func NewVeggiePizza() Pizza {
	this := new(VeggiePizza)
	this.name = "Veggie Pizza"
	this.dough = "Crust"
	this.sauce = "Marinara sauce"
	this.toppings = append(this.toppings, "Shredded mozzarella")
	this.toppings = append(this.toppings, "Grated parmesan")
	this.toppings = append(this.toppings, "Diced onion")
	this.toppings = append(this.toppings, "Sliced mushrooms")
	this.toppings = append(this.toppings, "Sliced red pepper")
	this.toppings = append(this.toppings, "Sliced black olives")
	return this
}
