package eg1

type Size = int8 //使用别名

const (
	TALL   Size = iota //小杯
	GRANDE             //中杯
	VENTI              //大杯
)

//饮料接口（装饰者和被装饰者都实现这个借口，这是关键）
type Beverage interface {
	GetDescription() string
	GetSize() Size
	SetSize(size Size)

	Cost() float64
}

//实现Beverage部分通用的方法（未实现Cost方法）
type BaseBeverage struct {
	description string
	size        Size
}

func (this *BaseBeverage) GetDescription() string {
	return this.description
}

func (this *BaseBeverage) GetSize() Size {
	return this.size
}

func (this *BaseBeverage) SetSize(size Size) {
	this.size = size
}

/*-------------------------------------------------饮料（实现Beverage接口）------------------------------------------*/
/*
* 1. 饮料1：浓缩咖啡
 */
type Espresso struct {
	BaseBeverage
}

func NewEspresso() *Espresso {
	this := new(Espresso)
	this.description = "Espresso"
	return this
}

func (this *Espresso) Cost() float64 {
	return 1.99
}

/*
* 2.饮料2：烘焙咖啡
 */
type DarkRoast struct {
	BaseBeverage
}

func NewDarkRoast() *DarkRoast {
	this := new(DarkRoast)
	this.description = "Dark Roast Coffee"
	return this
}

func (this *DarkRoast) Cost() float64 {
	return .99
}

/*
* 3. 饮料3
 */
type HouseBlend struct {
	BaseBeverage
}

func NewHouseBlend() *HouseBlend {
	this := new(HouseBlend)
	this.description = "House Blend Coffee"
	return this
}

func (this *HouseBlend) Cost() float64 {
	return .89
}
