package simple_barista

import (
	"testing"
	"fmt"
)

func TestBarista(t *testing.T)  {
	var tea CaffeineBeverage = NewTea()
	var coffee CaffeineBeverage = NewCoffee()

	fmt.Println("-----Making tea----")
	tea.PrepareRecipe()


	fmt.Println("----Making coffee----")
	coffee.PrepareRecipe()
}