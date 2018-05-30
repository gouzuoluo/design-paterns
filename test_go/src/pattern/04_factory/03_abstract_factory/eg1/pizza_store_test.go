package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {

	//1.在纽约的pizza店
	var nyStore PizzaStore = NewNYPizzaStore()

	var pizza Pizza = nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a ", pizza)

	pizza = nyStore.OrderPizza("clam")
	fmt.Println("Ethan ordered a ", pizza)

	//2.在芝加哥的pizza店
	var chicagoStore PizzaStore = NewChicagoPizzaStore()

	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a ", pizza)

	pizza = chicagoStore.OrderPizza("clam")
	fmt.Println("Joel ordered a ", pizza)
}
