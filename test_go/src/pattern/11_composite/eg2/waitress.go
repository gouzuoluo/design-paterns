package eg2

import (
	"fmt"
)

//女服务员
type Waitress struct {
	allMenus MenuComponent
}

func NewWaitress(allMenus MenuComponent) *Waitress {
	this := new(Waitress)
	this.allMenus = allMenus
	return this
}

//打印菜单
func (this *Waitress) PrintMenu() {
	this.allMenus.Print()
}

//打印素菜单(使用迭代器)
func (this *Waitress) PrintVegetarianMenu() {
	var iterator Iterator = this.allMenus.CreateIterator()
	fmt.Println("\n-------------\n打印素菜单\n-------------")

	for iterator.HasNext() {
		var menuComponent MenuComponent = iterator.Next()

		this.Try(
			func() {
				if menuComponent.IsVegetarian() {
					menuComponent.Print()
				}
			},
			func(arg interface{}) {
				//只是捕获异常，什么也不干
			})
	}
}

//go实现类似try...catch...的异常处理机制
func (*Waitress) Try(fun func(), handler func(interface{})) {
	defer func() {
		if ex := recover(); ex != nil {
			handler(ex)
		}
	}()

	fun()
}
