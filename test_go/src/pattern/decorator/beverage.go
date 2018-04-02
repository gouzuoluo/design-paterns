package decorator

type Size int8

const (
	TALL   Size = iota //小杯
	GRANDE             //中杯
	VENTI              //大杯
)

type Beverage interface {
	GetDescription() string
	GetSize() Size
	SetSize(size Size)
	Cost() float64
}

//实现Beverage部分方法
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

//未实现Cost方法

/*---------------------------------------------------饮料---------------------------------------------------------*/
//饮料1
type Espresso struct {
	BaseBeverage
}

func NewEspresso() Beverage {
	this := new(Espresso)
	this.description = "Espresso"
	return this
}

func (this *Espresso) Cost() float64 {
	return 1.99
}

//饮料2
type DarkRoast struct {
	BaseBeverage
}

func NewDarkRoast() Beverage {
	this := new(DarkRoast)
	this.description = "Dark Roast Coffee"
	return this
}

func (this *DarkRoast) Cost() float64 {
	return .99
}

//饮料3
type HouseBlend struct {
	BaseBeverage
}

func NewHouseBlend() Beverage {
	this := new(HouseBlend)
	this.description = "House Blend Coffee"
	return this
}

func (this *HouseBlend) Cost() float64 {
	return .89
}
