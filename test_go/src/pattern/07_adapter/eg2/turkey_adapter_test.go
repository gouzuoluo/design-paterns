package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var duck *MallardDuck = new(MallardDuck) //创建一只鸭子
	var turkey *WildTurkey = new(WildTurkey) //创建一只火鸡

	//测试火鸡，让它咯咯叫
	fmt.Println("-------The Turkey says--------")
	turkey.Gobble()

	//测试鸭子，让它呱呱加
	fmt.Println("-------The Duck says-------")
	duck.Quack()

	//火鸡适配器（火鸡假装是鸭子）
	//将火鸡包装进一个火鸡适配器中，使它看起来像一只鸭子
	var turkeyAdapter Duck = NewTurkeyAdapter(turkey)
	fmt.Println("-------The TurkeyAdapter says-------")
	turkeyAdapter.Quack()
}
