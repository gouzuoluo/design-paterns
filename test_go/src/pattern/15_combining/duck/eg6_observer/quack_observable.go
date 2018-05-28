package eg6_observer

//被观察者
//任何想被观察的Quackable都必须实现该接口；
//我们需要确定所有的Quackable都实现此接口，让Quackable接口继承该接口
type QuackObservable interface {
	RegisterObserver(Observer)
	NotifyObserver()
}

/*=================================================================================================================*/

//实现QuackObservable接口；
//我们只要把它插进一个类，就可以让该类将工作委托给Observable；
type Observable struct {
	duck      QuackObservable
	observers []Observer
}

func NewObservable(duck QuackObservable) *Observable {
	this := new(Observable)
	this.duck = duck
	this.observers = make([]Observer, 0, 10)

	return this
}

func (this *Observable) RegisterObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *Observable) NotifyObserver() {
	for _, observer := range this.observers {
		observer.Update(this.duck)
	}
}
