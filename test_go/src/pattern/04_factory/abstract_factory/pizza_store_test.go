package abstract_factory

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var nyStore PizzaStore = NewNYPizzaStore()
	var chicagoStore PizzaStore = NewChicagoPizzaStore()

	var pizza Pizza = nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a ", pizza)

	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a ", pizza)

	pizza = nyStore.OrderPizza("clam")
	fmt.Println("Ethan ordered a ", pizza)

	pizza = chicagoStore.OrderPizza("clam")
	fmt.Println("Joel ordered a ", pizza)
}
