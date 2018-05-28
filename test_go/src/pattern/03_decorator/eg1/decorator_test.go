package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var beverage Beverage = NewEspresso()
	beverage = NewMilk(beverage) //用牛奶来装饰浓缩咖啡，且返回类型也实现了新料接口
	fmt.Println(beverage.GetDescription() + " $" + fmt.Sprintf("%.2f", beverage.Cost()))

	var beverage1 Beverage = NewDarkRoast()
	beverage1 = NewMocha(beverage1)
	beverage1 = NewMocha(beverage1)
	beverage1 = NewWhip(beverage1) //多层次的装饰烘焙咖啡
	fmt.Println(beverage1.GetDescription()+" $", fmt.Sprintf("%.2f", beverage1.Cost()))

	var beverage2 Beverage = NewHouseBlend()
	beverage2.SetSize(VENTI)
	beverage2 = NewSoy(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)
	fmt.Println(beverage2.GetDescription()+" $", fmt.Sprintf("%.2f", beverage2.Cost()))
}
