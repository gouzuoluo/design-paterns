package eg2

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	api1 := NewAPI(1)
	s1 := api1.Say("Tom")
	fmt.Println(s1)

	api2 := NewAPI(2)
	s2 := api2.Say("Tom")
	fmt.Println(s2)
}
