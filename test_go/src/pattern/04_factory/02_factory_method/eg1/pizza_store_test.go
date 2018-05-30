package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {

	//在纽约店点pizza
	var nyStore PizzaStore = NewNYPizzaStore()
	var pizza Pizza = nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a",pizza.GetName())

	//在芝加哥店点pizza
	var chicagoStore PizzaStore = NewChicagoPizzaStore()
	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a",pizza.GetName())
}
