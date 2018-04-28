package example

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var duck *MallardDuck = new(MallardDuck)//创建一只鸭子
	var turkey *WildTurkey = new(WildTurkey)//创建一只火鸡

	var turkeyAdapter Duck = NewTurkeyAdapter(turkey)//将火鸡包装进一个火鸡适配器中，使它看起来像一只鸭子

	//测试这只火鸡，让它咯咯叫，让它飞行
	fmt.Println("-------The Turkey says--------")
	turkey.Gobble()
	turkey.Fly()

	//使用testDuck测试鸭子，需要传入一个鸭子对象
	fmt.Println("-------The Duck says-------")
	testDuck(duck)

	//测试传入假装是鸭子的火鸡
	fmt.Println("-------The TurkeyAdapter says-------")
	testDuck(turkeyAdapter)
}

func testDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}