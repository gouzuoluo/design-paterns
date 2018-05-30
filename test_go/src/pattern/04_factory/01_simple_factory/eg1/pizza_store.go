package eg1

//简单工厂（通过一个单独的类来创建对象）
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

//披萨店（工厂的客户）
type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore(factory *SimplePizzaFactory) *PizzaStore {
	this := new(PizzaStore)
	this.factory = factory
	return this
}

//点单
func (this *PizzaStore) OrderPizza(tp string) Pizza {

	pizza := this.factory.CreatePizza(tp) //把new操作替换成工厂对象的创建方法。

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
