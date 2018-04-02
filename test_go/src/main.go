package main

import "fmt"

type A interface {
	getA()
}

type B interface {
	A
	getB()
}

type C struct {
	B
}


func (this *C)getA()  {
	fmt.Println("AAAAAAAAA")
}

func (this *C)getB()  {
	fmt.Println("BBBBBBBBB")
}

func main() {
	var a B = new(C)
	a.getA()
	a.getB()
}
