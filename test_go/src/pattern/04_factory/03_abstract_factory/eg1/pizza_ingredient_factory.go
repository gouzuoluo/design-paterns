package eg1

/*
*   披萨原料工厂
*/

//原料工厂接口（抽象工厂）
type PizzaIngredientFactory interface {
	CreateDough() Dough         //创建面团
	CreateSauce() Sauce         //创建酱油
	CreateCheese() Cheese       //创建奶酪
	CreateVeggies() []Veggies   //创建蔬菜
	CreatePepperoni() Pepperoni //创建意大利香肠
	CreateClam() Clams          //创建蛤蜊
}

/*====================================================================================================================*/

/*
* 1.纽约原料工厂
 */
type NYPizzaIngredientFactory struct {
}

func (this *NYPizzaIngredientFactory) CreateDough() Dough {
	return new(ThinCrustDough)
}

func (this *NYPizzaIngredientFactory) CreateSauce() Sauce {
	return new(MarinaraSauce)
}

func (this *NYPizzaIngredientFactory) CreateCheese() Cheese {
	return new(ReggianoCheese)
}

func (this *NYPizzaIngredientFactory) CreateVeggies() []Veggies {
	var veggige []Veggies = []Veggies{new(Garlic), new(Onion), new(Mushroom)}
	return veggige
}

func (this *NYPizzaIngredientFactory) CreatePepperoni() Pepperoni {
	return new(SlicedPepperoni)
}

func (this *NYPizzaIngredientFactory) CreateClam() Clams {
	return new(FreshClams)
}

/*
* 2.芝加哥原料工厂
 */
type ChicagoPizzaIngredientFactory struct {
}

func (this *ChicagoPizzaIngredientFactory) CreateDough() Dough {
	return new(ThickCrustDough)
}

func (this *ChicagoPizzaIngredientFactory) CreateSauce() Sauce {
	return new(PlumTomatoSauce)
}

func (this *ChicagoPizzaIngredientFactory) CreateCheese() Cheese {
	return new(MozzarellaCheese)
}

func (this *ChicagoPizzaIngredientFactory) CreateVeggies() []Veggies {
	var veggies []Veggies = []Veggies{new(BlackOlives), new(Spinach), new(Eggplant)}
	return veggies
}

func (this *ChicagoPizzaIngredientFactory) CreatePepperoni() Pepperoni {
	return new(SlicedPepperoni)
}

func (this *ChicagoPizzaIngredientFactory) CreateClam() Clams {
	return new(FrozenClams)
}
