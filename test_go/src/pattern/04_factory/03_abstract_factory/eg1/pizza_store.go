package eg1

/*
*  披萨店
*/

import "fmt"

type PizzaStore interface {
	OrderPizza(tp string) Pizza

	CreatePizza(tp string) Pizza //这里是一个工厂方法
}

//使用匿名包含接口（未实现CreatePizza方法）
type basePizzaStore struct {
	PizzaStore
}

func (this *basePizzaStore) OrderPizza(tp string) Pizza {
	pizza := this.PizzaStore.CreatePizza(tp)

	fmt.Println("--- Making a ", pizza.GetName(), " ---")

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

/*====================================================================================================================*/

/*
* 1.纽约披萨店
 */
type NYPizzaStore struct {
	basePizzaStore
}

func NewNYPizzaStore() PizzaStore {
	this := new(NYPizzaStore)
	this.PizzaStore = this
	return this
}

func (this *NYPizzaStore) CreatePizza(tp string) Pizza {
	var pizza Pizza = nil
	var ingredientFactory PizzaIngredientFactory = new(NYPizzaIngredientFactory) //这里是一个抽象工厂
	switch tp {
	case "cheese":
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("New York Style Cheese Pizza")
	case "veggies":
		pizza = NewVeggiePizza(ingredientFactory)
		pizza.SetName("New York Style Veggies Pizza")
	case "clam":
		pizza = NewClamPizza(ingredientFactory)
		pizza.SetName("New York Style Clam Pizza")
	case "pepperoni":
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("New York Style Pepperoni Pizza")
	}

	return pizza
}

/*
* 2.芝加哥披萨店
 */
type ChicagoPizzaStore struct {
	basePizzaStore
}

func NewChicagoPizzaStore() PizzaStore {
	this := new(ChicagoPizzaStore)
	this.PizzaStore = this
	return this
}

func (this *ChicagoPizzaStore) CreatePizza(tp string) Pizza {
	var pizza Pizza = nil
	var ingredientFactory PizzaIngredientFactory = new(ChicagoPizzaIngredientFactory)
	switch tp {
	case "cheese":
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("Chicago Style Cheese Pizza")
	case "veggies":
		pizza = NewVeggiePizza(ingredientFactory)
		pizza.SetName("Chicago Style Veggie Pizza")
	case "clam":
		pizza = NewClamPizza(ingredientFactory)
		pizza.SetName("Chicago Style Clam Pizza")
	case "pepperoni":
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("Chicago Style Pepperoni Pizza")
	}

	return pizza
}
