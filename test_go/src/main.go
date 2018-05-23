package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Unix(0, 0))
}