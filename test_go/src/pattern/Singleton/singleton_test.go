package Singleton

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	s := GetSingletonInstance()
	fmt.Println(s)
}
