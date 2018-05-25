package factory_method

type PizzaStore interface {
	OrderPizza(tp string) Pizza
	CreatePizza(tp string) Pizza
}

//实现PizzaStore接口
type BasePizzaStore struct {
	PizzaStore
}

func (this *BasePizzaStore) OrderPizza(tp string) Pizza {
	pizza := this.PizzaStore.CreatePizza(tp) //注意这里
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

/*-------------------------------------------------------------------------------------------------------------------*/
//1.纽约披萨店
type NYPizzaStore struct {
	BasePizzaStore
}

func NewNYPizzaStore() PizzaStore {
	this := new(NYPizzaStore)
	this.PizzaStore = this //注意这里
	return this
}

func (this *NYPizzaStore) CreatePizza(tp string) Pizza {
	if tp == "cheese" {
		return NewNYStyleCheesePizza()
	} else if tp == "veggie" {
		return NewNYStyleVeggiePizza()
	} else if tp == "clam" {
		return NewNYStyleClamPizza()
	} else if tp == "pepperoni" {
		return NewNYStylePepperoniPizza()
	} else {
		return nil
	}
}

//2.芝加哥披萨店
type ChicagoPizzaStore struct {
	BasePizzaStore
}

func NewChicagoPizzaStore() PizzaStore {
	this := new(ChicagoPizzaStore)
	this.PizzaStore = this //注意这里
	return this
}

func (this *ChicagoPizzaStore) CreatePizza(tp string) Pizza {
	if tp == "cheese" {
		return NewChicagoStyleCheesePizza()
	} else if tp == "veggie" {
		return NewChicagoStyleVeggiePizza()
	} else if tp == "clam" {
		return NewChicagoStyleClamPizza()
	} else if tp == "pepperoni" {
		return NewChicagoStylePepperoniPizza()
	} else {
		return nil
	}
}
