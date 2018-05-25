package simple_factory

//简单工厂，通过一个单独的类来创建对象
type SimplePizzaFactory struct {
}

func (this *SimplePizzaFactory) CreatePizza(tp string) Pizza {
	var pizza Pizza = nil
	if tp == "cheese" {
		pizza = NewCheesePizza()
	} else if tp == "pepperoni" {
		pizza = NewPepperoniPizza()
	} else if tp == "clam" {
		pizza = NewClamPizza()
	} else if tp == "veggie" {
		pizza = NewVeggiePizza()
	}
	return pizza
}

/*--------------------------------------------------------------------------------------------------------------*/
type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore(factory *SimplePizzaFactory) *PizzaStore {
	this := new(PizzaStore)
	this.factory = factory
	return this
}

func (this *PizzaStore) OrderPizza(tp string) Pizza {
	pizza := this.factory.CreatePizza(tp)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
