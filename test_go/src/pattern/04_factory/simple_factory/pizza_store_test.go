package simple_factory

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	store := NewPizzaStore(new(SimplePizzaFactory))

	var pizza Pizza = store.OrderPizza("cheese")
	fmt.Println("We ordered a", pizza.GetName())

	pizza = store.OrderPizza("veggie")
	fmt.Println("We ordered a", pizza.GetName())

	fmt.Println(pizza)
}
