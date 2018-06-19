package eg3

/*
* 外部迭代器
*/

import "fmt"

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}

/*===================================================================================================================*/

//数字集合迭代器
type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func NewNumbersIterator(numbers *Numbers, start int) Iterator {
	this := new(NumbersIterator)
	this.numbers = numbers
	this.next = start
	return this
}

func (this *NumbersIterator) First() {
	this.next = this.numbers.start
}

func (this *NumbersIterator) IsDone() bool {
	return this.next > this.numbers.end
}

func (this *NumbersIterator) Next() interface{} {
	if !this.IsDone() {
		next := this.next
		this.next++
		return next
	}
	return nil
}
