package simple_romote

type Command interface {
	Execute()
}

/*---------------------------------------------ConcreteCommand--------------------------------*/
//1.开灯命令
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	this := new(LightOnCommand)
	this.light = light
	return this
}

func (this *LightOnCommand) Execute() {
	this.light.On()
}

//2.关灯命令
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	this := new(LightOffCommand)
	this.light = light
	return this
}

func (this *LightOffCommand) Execute() {
	this.light.Off()
}

//3.开车库门命令
type GarageDoorOpenCommand struct {
	door *GarageDoor
}

func NewGarageDoorOpenCommand(door *GarageDoor) *GarageDoorOpenCommand {
	this := new(GarageDoorOpenCommand)
	this.door = door
	return this
}

func (this *GarageDoorOpenCommand) Execute() {
	this.door.Up()
}
