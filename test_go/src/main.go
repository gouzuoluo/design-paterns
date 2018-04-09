package main

import "fmt"

type A interface {
	f1()
	f2()
}

type bA struct {
}

func (this *bA)f1()  {
	fmt.Println("bA----f1")
	this.f2()
}

func (this *bA)f2()  {
	fmt.Println("bA-----f2")
}

/*--------------------------------------------------------------------------------*/
//继承bA
type A1 struct {
	bA
}
//重写f2
func (this *A1)f2()  {
	fmt.Println("A1------f2")
	this.bA.f2()
}

func main() {
	var a A = new(A1)

	a.f1()
	a.f2()
}
