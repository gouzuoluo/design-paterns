package main

import "fmt"

func main() {
	ts := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("---",ts)
	for i, v := range ts {
		if v == 2 {
			ts = append(ts[:i], ts[i+1:]...)
		}
	}

	fmt.Println("***",ts)
}
