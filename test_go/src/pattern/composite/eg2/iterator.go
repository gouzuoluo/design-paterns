package eg2

import (
	"utils/container/stack"
)

//迭代器接口
type Iterator interface {
	HasNext() bool              //是否还有下一个菜单项
	Next() MenuComponent        //获取下一个菜单组件
}

/*====================================================================================================================*/
/*
* 1.空迭代器
*/
type NilIterator struct {
}

func NewNilIterator() Iterator {
	this := new(NilIterator)
	return this
}

func (this *NilIterator)Next() MenuComponent {
	return nil
}

func (this *NilIterator)HasNext() bool {
	return false
}

/*-------------------------------------------------------------------------------------------------------------------*/


/*
* 2.菜单集合迭代器
*/
type MenuGatherIterator struct {
	menuComponents []MenuComponent
	position       int
}

func NewMenuGatherIterator(menuComponents []MenuComponent) Iterator {
	this := new(MenuGatherIterator)
	this.menuComponents = menuComponents
	this.position = 0
	return this
}

func (this *MenuGatherIterator)Next() MenuComponent {
	if this.HasNext() {
		menuComponent := this.menuComponents[this.position]
		this.position += 1
		return menuComponent
	} else {
		return nil
	}
}

func (this *MenuGatherIterator)HasNext() bool {
	if len(this.menuComponents) > this.position {
		return true
	} else {
		return false
	}
}

/*
* 3.组件迭代器
*/
type CompositeIterator struct {
	stack *stack.Stack //Element是MenuGatherIterator类型(即是Iterator)
}

func NewCompositeIterator(iterator Iterator) Iterator {
	this := new(CompositeIterator)
	this.stack = stack.NewStack()
	this.stack.Push(iterator)
	return this
}

func (this *CompositeIterator)Next() MenuComponent {
	if this.HasNext() {
		iterator := this.stack.Top().(Iterator)
		component := iterator.Next()
		this.stack.Push(component.CreateIterator())
		return component
	} else {
		return nil
	}
}

func (this *CompositeIterator)HasNext() bool {
	if this.stack.Empty() {
		return false
	} else {
		iterator := this.stack.Top().(Iterator)
		if !iterator.HasNext() {
			this.stack.Pop()
			return this.HasNext()
		} else {
			return true
		}
	}
}
