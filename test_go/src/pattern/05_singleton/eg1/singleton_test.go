package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	s := GetSingletonInstance()
	fmt.Println(s)
}
