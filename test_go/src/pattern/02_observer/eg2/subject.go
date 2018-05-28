package eg2

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	this := new(Subject)
	this.observers = make([]Observer, 0, 4)
	return this
}

func (this *Subject) Attach(o Observer) {
	this.observers = append(this.observers, o)
}

func (this *Subject) notify() {
	for _, o := range this.observers {
		o.Update(this)
	}
}

func (this *Subject) UpdateContext(context string) {
	this.context = context
	this.notify()
}
