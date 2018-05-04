package example1

import "container/list"

//迭代器接口
type Iterator interface {
	HasNext() bool//是否还有下一个菜单项
	Next() *MenuItem//获取下一个菜单项
}

/*-------------------------------------------------------------------------------------------------------------------*/

//1.煎饼屋菜单(早餐菜单)的迭代器
type PancakeHouseMenuIterator struct {
	items    *list.List
	position *list.Element
}

func NewPancakeHouseMenuIterator(items *list.List) Iterator {
	this := new(PancakeHouseMenuIterator)
	this.items = items
	this.position = this.items.Front()
	return this
}

//实现Iterator接口的Next方法
func (this *PancakeHouseMenuIterator)Next() *MenuItem {
	if this.position != nil {
		item := this.position.Value.(*MenuItem)
		this.position = this.position.Next()
		return item
	} else {
		return nil
	}
}

//实现Iterator接口的HasNext方法
func (this *PancakeHouseMenuIterator)HasNext() bool {
	if this.position.Next() != nil {
		return true
	} else {
		return false
	}
}

//2.餐厅菜单（午餐菜单）的迭代器
type DinerMenuIterator struct {
	menuItems []*MenuItem
	position  int
}

func NewDinerMenuIterator(menuItems []*MenuItem) Iterator {
	this := new(DinerMenuIterator)
	this.menuItems = menuItems
	this.position = 0
	return this
}

//实现Iterator接口的Next方法
func (this *DinerMenuIterator)Next() *MenuItem {
	if len(this.menuItems) > this.position {
		menuItem := this.menuItems[this.position]
		this.position += 1
		return menuItem
	} else {
		return nil
	}
}

//实现Iterator接口的hasNext方法
func (this *DinerMenuIterator)HasNext() bool {
	if len(this.menuItems) > this.position {
		return true
	} else {
		return false
	}
}