package factory_method

import (
	"fmt"
	"bytes"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName()string
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

func (this *BasePizza)GetName()string  {
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
/*---------------------------------------------纽约风味披萨-----------------------------------------------------------*/
//1奶酪披萨
type NYStyleCheesePizza struct {
	BasePizza
}

func NewNYStyleCheesePizza() Pizza {
	this := new(NYStyleCheesePizza)
	this.name = "NY Style Sauce and Cheese Pizza"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	return this
}

//2.蛤蜊披萨
type NYStyleClamPizza struct {
	BasePizza
}

func NewNYStyleClamPizza() Pizza {
	this := new(NYStyleClamPizza)
	this.name = "NY Style Clam Pizzaa"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	this.toppings = append(this.toppings, "Fresh Clams from Long Island Sound")
	return this
}

//3.意大利辣香肠披萨
type NYStylePepperoniPizza struct {
	BasePizza
}

func NewNYStylePepperoniPizza() Pizza {
	this := new(NYStylePepperoniPizza)
	this.name = "NY Style Pepperoni Pizz"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	this.toppings = append(this.toppings, "Sliced Pepperoni")
	this.toppings = append(this.toppings, "Garlic")
	this.toppings = append(this.toppings, "Onion")
	this.toppings = append(this.toppings, "Mushrooms")
	this.toppings = append(this.toppings, "Red Pepper")
	return this
}

//4.蔬菜披萨
type NYStyleVeggiePizza struct {
	BasePizza
}

func NewNYStyleVeggiePizza() Pizza {
	this := new(NYStyleVeggiePizza)
	this.name = "NY Style Veggie Pizza"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings,"Grated Reggiano Cheese")
	this.toppings = append(this.toppings,"Garlic")
	this.toppings = append(this.toppings,"Onion")
	this.toppings = append(this.toppings,"Mushrooms")
	this.toppings = append(this.toppings,"Red Pepper")
	return this
}


/*---------------------------------------------芝加哥风味披萨-----------------------------------------------------------*/
//1奶酪披萨
type ChicagoStyleCheesePizza struct {
	BasePizza
}

func NewChicagoStyleCheesePizza() Pizza {
	this := new(ChicagoStyleCheesePizza)
	this.name = "Chicago Style Sauce and Cheese Pizza"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	return this
}

//2.蛤蜊披萨
type ChicagoStyleClamPizza struct {
	BasePizza
}

func NewChicagoStyleClamPizza() Pizza {
	this := new(ChicagoStyleClamPizza)
	this.name = "Chicago Style Clam Pizzaa"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	this.toppings = append(this.toppings, "Fresh Clams from Long Island Sound")
	return this
}

//3.意大利辣香肠披萨
type ChicagoStylePepperoniPizza struct {
	BasePizza
}

func NewChicagoStylePepperoniPizza() Pizza {
	this := new(ChicagoStylePepperoniPizza)
	this.name = "Chicago Style Pepperoni Pizz"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara sauce"
	this.toppings = append(this.toppings, "Grated Reggiano Cheese")
	this.toppings = append(this.toppings, "Sliced Pepperoni")
	this.toppings = append(this.toppings, "Garlic")
	this.toppings = append(this.toppings, "Onion")
	this.toppings = append(this.toppings, "Mushrooms")
	this.toppings = append(this.toppings, "Red Pepper")
	return this
}

//4.蔬菜披萨
type ChicagoStyleVeggiePizza struct {
	BasePizza
}

func NewChicagoStyleVeggiePizza() Pizza {
	this := new(ChicagoStyleVeggiePizza)
	this.name = "Chicago Style Veggie Pizza"
	this.dough = "Thin Crust Dough"
	this.sauce = "Marinara Sauce"
	this.toppings = append(this.toppings,"Grated Reggiano Cheese")
	this.toppings = append(this.toppings,"Garlic")
	this.toppings = append(this.toppings,"Onion")
	this.toppings = append(this.toppings,"Mushrooms")
	this.toppings = append(this.toppings,"Red Pepper")
	return this
}
