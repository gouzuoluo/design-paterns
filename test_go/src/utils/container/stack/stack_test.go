package stack

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	if stack.Empty() {
		fmt.Println("Stack为空! ")
	} else {
		fmt.Println("Stack不为空! ", stack.Size())
	}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println("当前Size() = ", stack.Size())
	stack.Print()

	fmt.Println("当前Top() = ", stack.Top())

	stack.Pop()
	fmt.Println("执行完Pop()后的Top() = ", stack.Top())
	stack.Print()

	stack.Set(2, 900)
	fmt.Println("\n执行完Set(2,900)后的Stack")
	stack.Print()

	fmt.Println("\nGet()查看指定的元素: ")
	fmt.Println("当前idx为1的元素 = ", stack.Get(1))
	fmt.Println("当前idx为2的元素 = ", stack.Get(2))

	stack2 := NewStack()
	stack2.Push("111")
	stack2.Push("222")
	fmt.Println("\nstack2的初始内容:")
	stack2.Print()

	stack.Swap(stack2)
	fmt.Println("Swap()后stack的内容:")
	stack.Print()
	fmt.Println("Swap()后stack2的内容:")
	stack2.Print()

	fmt.Println("\nstack增加字符串元素: ")
	stack.Push("中文元素")
	stack.Push("elem1")
	stack.Print()

}
