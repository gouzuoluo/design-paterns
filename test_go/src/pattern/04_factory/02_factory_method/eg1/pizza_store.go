package eg1

/* ===================高层组件================*/

type PizzaStore interface {
	//点单Pizza（超类实现该方法）
	OrderPizza(tp string) Pizza

	//创建pizza（每个子类实现该方法。让子类做决定）
	CreatePizza(tp string) Pizza //这就是一个工厂方法，依赖子类来处理对象的创建
}

//实现PizzaStore接口OrderPizza方法（未实现CreatePizza方法）
type BasePizzaStore struct {
	PizzaStore
}

//点单
func (this *BasePizzaStore) OrderPizza(tp string) Pizza {

	pizza := this.PizzaStore.CreatePizza(tp) //工厂方法将客户（即这个OrderPizza方法）和实际创建具体产品的代码分离开来

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

/*====================================================================================================================*/

/*
* 1.纽约披萨店
 */
type NYPizzaStore struct {
	BasePizzaStore
}

func NewNYPizzaStore() PizzaStore {
	this := new(NYPizzaStore)
	this.PizzaStore = this //注意这里
	return this
}

//真正的工厂方法1
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

/*
* 2.芝加哥披萨店
 */
type ChicagoPizzaStore struct {
	BasePizzaStore
}

func NewChicagoPizzaStore() PizzaStore {
	this := new(ChicagoPizzaStore)
	this.PizzaStore = this //注意这里
	return this
}

//真正的工厂方法2
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
