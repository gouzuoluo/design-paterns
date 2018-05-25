/*
*  策略模式:定义了一族算法，分别封装起来，让他们之间可以互相替换，此模式让算法的
*  			的变化独立于使用算法的客户。
*/
package _1_strategy

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var duck Duck = NewMallardDuck()
	fmt.Println("----------------------------------------")
	duck.Display()
	duck.Swim()
	duck.PreformFly()
	duck.PreformQuack()

	duck.SetFlyBehavior(new(FlyRocketPowered))
	duck.PreformFly()

	fmt.Println("----------------------------------------")
	duck = NewModelDuck()
	duck.Display()
	duck.Swim()
	duck.PreformFly()
	duck.PreformQuack()

	duck.SetQuackBehavior(new(Squeak))
	duck.PreformQuack()
	fmt.Println("----------------------------------------")
}
