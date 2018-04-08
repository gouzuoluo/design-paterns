package main

import "fmt"

type Printer interface {
	Print()
}

type CanonPrinter struct {
	PrintName string
}

func (p CanonPrinter) Print() {
	fmt.Println("CanonPrinte")
}

type PrintWorker struct {
	Printer
	name string
	age  int
}

// 如果没有下面实现，则
//func (pw PrintWorker) Print() {
//	fmt.Println("PrintWorker")
//	pw.Printer.Print() // 这里 pw 首先引用内部嵌入Printer接口的实例，然后调用Printer 接口实例的Print()方法
//}

func main() {
	canon := CanonPrinter{"canoprint_num_1"}
	p := PrintWorker{Printer: canon, name: "ansendong", age: 34}
	p.Print() //如果没有上述红色部分的实现，则这里只调用CanonPrinter实现的Print()方法。
}
