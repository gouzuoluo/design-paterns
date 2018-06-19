package eg1

//女服务员
type Waitress struct {
	allMenus MenuComponent
}

func NewWaitress(allMenus MenuComponent) *Waitress {
	this := new(Waitress)
	this.allMenus = allMenus
	return this
}

//打印菜单
func (this *Waitress) PrintMenu() {
	this.allMenus.Print()
}
