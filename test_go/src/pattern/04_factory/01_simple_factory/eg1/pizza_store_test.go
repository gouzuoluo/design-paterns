package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	store := NewPizzaStore(new(SimplePizzaFactory))

	//1.点单cheese
	var pizza Pizza = store.OrderPizza("cheese")
	fmt.Println("We ordered a", pizza.GetName())

	//2.点单veggie
	pizza = store.OrderPizza("veggie")
	fmt.Println("We ordered a", pizza.GetName())

	fmt.Println(pizza)
}
