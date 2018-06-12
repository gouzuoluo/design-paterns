package eg4

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (this *Box) PressButton1() {
	this.button1.Execute()
}

func (this *Box) PressButton2() {
	this.button2.Execute()
}
