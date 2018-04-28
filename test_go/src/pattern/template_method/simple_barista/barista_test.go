package simple_barista

import (
	"testing"
	"fmt"
)

func TestBarista(t *testing.T)  {
	var tea Tea = new(Tea)
	var coffee Coffee = new(Coffee)

	fmt.Println("Making tea...")
	tea.PrepareRecipe()


	fmt.Println("Making coffee...")
	coffee.PrepareRecipe()
}