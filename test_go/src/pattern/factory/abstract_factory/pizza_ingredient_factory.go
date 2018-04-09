package abstract_factory


type PizzaIngredientFactory interface {
	CreateDough() Dough         //创建面团
	CreateSauce() Sauce         //创建酱油
	CreateCheese() Cheese       //创建奶酪
	CreateVeggies() []Veggies   //创建蔬菜
	CreatePepperoni() Pepperoni //创建意大利香肠
	CreateClam() Clams          //创建蛤蜊
}

/*------------------------------------------------------------------------------------------------------------------*/
//1.纽约原料工厂