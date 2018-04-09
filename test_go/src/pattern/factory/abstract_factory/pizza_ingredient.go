package abstract_factory

//生面团
type Dough interface {
	String() string
}

//酱油
type Sauce interface {
	String() string
}

//蔬菜
type Veggies interface {
	String() string
}

//奶酪
type Cheese interface {
	String() string
}


//意大利香肠
type Pepperoni interface {
	String() string
}

//蛤蜊
type Clams interface {
	String() string
}