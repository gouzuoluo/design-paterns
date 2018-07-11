package eg3

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	this := new(DayContext)
	this.today = new(Sunday)
	return this
}

func (this *DayContext) Today() {
	this.today.Today()
}

func (this *DayContext) Next() {
	this.today.Next(this)
}
