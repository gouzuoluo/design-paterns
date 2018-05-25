package abstract_factory

/*
* 披萨原料
*/


//生面团
type Dough interface {
	String() string
}

//1.薄皮面
type ThinCrustDough struct {
}

func (this *ThinCrustDough) String() string {
	return "Thin Crust Dough"
}

//2.厚皮面
type ThickCrustDough struct {
}

func (this *ThickCrustDough) String() string {
	return "ThickCrust style extra thick crust dough"
}

/*-------------------------------------------------------*/

//酱油
type Sauce interface {
	String() string
}

//1.海员沙司酱油
type MarinaraSauce struct {
}

func (this *MarinaraSauce) String() string {
	return "Marinara Sauce"
}

//2.番茄酱
type PlumTomatoSauce struct {
}

func (this *PlumTomatoSauce) String() string {
	return "Tomato sauce with plum tomatoes"
}

/*-------------------------------------------------------*/

//蔬菜
type Veggies interface {
	String() string
}

//1.红椒
type RedPepper struct {
}

func (this *RedPepper) String() string {
	return "Red Pepper"
}

//2.大蒜
type Garlic struct {
}

func (this *Garlic) String() string {
	return "Garlic"
}

//3.洋葱
type Onion struct {
}

func (this *Onion) String() string {
	return "Onion"
}

//4.蘑菇
type Mushroom struct {
}

func (this *Mushroom) String() string {
	return "Mushrooms"
}

//5.黑橄榄
type BlackOlives struct {
}

func (this *BlackOlives) String() string {
	return "Black Olives"
}

//6.菠菜
type Spinach struct {
}

func (this *Spinach) String() string {
	return "Spinach"
}

//7.茄子
type Eggplant struct {
}

func (this *Eggplant) String() string {
	return "Eggplant"
}

/*-------------------------------------------------------*/

//奶酪
type Cheese interface {
	String() string
}

//1.巴马干酪
type ReggianoCheese struct {
}

func (this *ReggianoCheese) String() string {
	return "Reggiano Cheese"
}

//2.意大利白干酪
type MozzarellaCheese struct {
}

func (this *MozzarellaCheese) String() string {
	return "Shredded Mozzarella"
}

//3.帕尔马干酪
type ParmesanCheese struct {

}

func (this *ParmesanCheese) String() string {
	return "Shredded Parmesan"
}
/*-------------------------------------------------------*/

//意大利香肠
type Pepperoni interface {
	String() string
}

//切片式意大利香肠
type SlicedPepperoni struct {
}

func (this *SlicedPepperoni) String() string {
	return "Sliced Pepperoni"
}

/*-------------------------------------------------------*/

//蛤蜊
type Clams interface {
	String() string
}

//1.新鲜蛤蜊
type FreshClams struct {
}

func (this *FreshClams) String() string {
	return "Fresh Clams from Long Island Sound"
}

//1.冰冻蛤蜊
type FrozenClams struct {
}

func (this *FrozenClams) String() string {
	return "Frozen Clams from Chesapeake Bay"
}

/*-------------------------------------------------------*/
