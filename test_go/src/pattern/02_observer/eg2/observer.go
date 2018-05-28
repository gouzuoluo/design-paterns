package eg2

import "fmt"

type Observer interface {
	Update(*Subject)
}

/*====================================================================================================================*/

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	this := new(Reader)
	this.name = name
	return this
}

func (this *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s \n", this.name, s.context)
}
