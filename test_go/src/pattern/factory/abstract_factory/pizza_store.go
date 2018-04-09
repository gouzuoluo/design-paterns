package abstract_factory

import "fmt"

type PizzaStore interface {
	OrderPizza(tp string) Pizza
	CreatePizza(tp string) Pizza
}

//使用匿名包含接口
type basePizzaStore struct {
	PizzaStore
}

//未实现CreatePizza方法

func (this *basePizzaStore) OrderPizza(tp string) Pizza {
	pizza := this.PizzaStore.CreatePizza(tp)
	fmt.Println("--- Making a ", pizza.GetName(), " ---")
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

/*-------------------------------------------------------------------------------------------------------------------*/
//1.纽约披萨店
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
	var ingredientFactory PizzaIngredientFactory = NewNYPizzaIngredientFactory()
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

//2.芝加哥披萨店
