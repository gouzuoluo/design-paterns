package eg3

//集合接口
type Aggregate interface {
	Iterator() Iterator
}

/*===================================================================================================================*/

//数字集合
type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (this *Numbers) Iterator() Iterator {
	return NewNumbersIterator(this, this.start)
}
