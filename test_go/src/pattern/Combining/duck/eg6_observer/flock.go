package eg6_observer

//群（实现Quackable接口）
//组合需要和叶节点元素实现相同的接口；这里的叶节点就是Quackable
type Flock struct {
	quackers []Quackable //这里最好使用带有迭代器的容器
}

func NewFlock() *Flock {
	this := new(Flock)
	this.quackers = make([]Quackable, 0, 10)

	return this
}

func (this *Flock) Add(quacker Quackable) {
	this.quackers = append(this.quackers, quacker)
}

func (this *Flock) Quack() {
	for _, quacker := range this.quackers {
		if v, ok := quacker.(*Flock); ok {
			v.Quack()
		} else {
			quacker.Quack()
		}
	}
}

func (this *Flock) RegisterObserver(observer Observer) {
	for _, quacker := range this.quackers {
		if v, ok := quacker.(*Flock); ok {
			v.RegisterObserver(observer)
		} else {
			quacker.RegisterObserver(observer)
		}
	}
}

func (this *Flock) NotifyObserver() {
}

func (this *Flock) String() string {
	return "Flock of Quackers"
}
