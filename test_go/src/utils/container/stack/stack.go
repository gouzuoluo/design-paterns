package stack

import "fmt"

type Stack struct {
	Element []interface{} //元素
}

func NewStack() *Stack {
	this := new(Stack)
	this.Element = make([]interface{}, 0, 10)
	return this
}

func (this *Stack)Size() int {
	return len(this.Element)
}

//是否为空
func (this *Stack)Empty() bool {
	if this.Element == nil || this.Size() == 0 {
		return true
	} else {
		return false
	}
}

//压栈（添加元素）
func (this *Stack)Push(value ...interface{}) {
	this.Element = append(this.Element, value...)
}

//返回栈顶元素
func (this *Stack)Top() interface{} {
	if this.Size() > 0 {
		return this.Element[this.Size() - 1]
	}
	return nil//read empty stack
}

//移除栈顶元素
func (this *Stack)Pop() error {
	if this.Size() > 0 {
		this.Element = this.Element[:this.Size() - 1]
		return nil
	} else {
		return fmt.Errorf("Stack为空.")
	}
}

//交换堆栈值
func (this *Stack)Swap(other *Stack) {
	switch  {
	case this.Size() == 0 && other.Size() == 0:
		return
	case other.Size() == 0:
		other.Element = this.Element
		this.Element = nil
	case this.Size() == 0:
		this.Element = other.Element
		other.Element = nil
	default:
		this.Element, other.Element = other.Element, this.Element
	}
	return
}

//修改指定索引元素
func (this *Stack)Set(idx int, value interface{}) error {
	if idx >= 0 && this.Size() > 0 && this.Size() > idx {
		this.Element[idx] = value
		return nil
	} else {
		return fmt.Errorf("Set失败!")
	}
}

//返回指定索引的元素
func (this *Stack)Get(idx int) interface{} {
	if idx >= 0 && this.Size() > 0 && this.Size() > idx {
		return this.Element[idx]
	} else {
		return nil
	}
}

//打印
func (this *Stack)Print() {
	for i, element := range this.Element {
		fmt.Println(i, "=>", element)
	}
}

