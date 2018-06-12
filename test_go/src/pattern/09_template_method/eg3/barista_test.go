package simple_barista

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T)  {
	var tea CaffeineBeverageWithHook = NewTeaWithHook()
	var coffee CaffeineBeverageWithHook = NewCoffeeWithHook()

	fmt.Println("---Hook--Making tea----")
	tea.PrepareRecipe()


	fmt.Println("--Hook--Making coffee----")
	coffee.PrepareRecipe()
}