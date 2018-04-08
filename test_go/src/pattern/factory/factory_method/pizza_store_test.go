package factory_method

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var nyStore PizzaStore = NewNYPizzaStore()
	var chicagoStore PizzaStore = NewChicagoPizzaStore()

	var pizza Pizza = nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a",pizza.GetName())

	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a",pizza.GetName())
}
