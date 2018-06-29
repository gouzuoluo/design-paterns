package main

import (
	"fmt"
)

func main() {
	ff("slots_freeone.CFG",
		"slots_freetwo.CFG", "slots_freethree.CFG",
		"slots_freefour.CFG", "slots_freefive.CFG")
}

func ff(s ...string) {
	var ts []string
	if len(s) > 0 {
		for _, v := range s {
			fmt.Println("---", v)
		}
		fmt.Println("===", s)

		ts = s
	}

	fmt.Println("****", ts)
}
