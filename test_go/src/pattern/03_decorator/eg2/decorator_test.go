package eg2

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var c Component = NewConcreteComponent()
	c = NewAddDecorator(c, 10)
	c = NewMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d \n", res)
}
