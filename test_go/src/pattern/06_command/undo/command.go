package romote

type Command interface {
	Execute()
	Undo()
}

/*--------------------------------------------------------------------------------------*/
//0.无具体操作命令
type NoCommand struct {
}

func (this *NoCommand) Execute() {
	//什么也不干
}

func (this *NoCommand) Undo() {
	//什么也不干
}

//1.关闭吊扇
type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int //记录吊扇之前速度
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) *CeilingFanOffCommand {
	this := new(CeilingFanOffCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanOffCommand) Execute() {
	this.prevSpeed = this.ceilingFan.GetSpeed()
	this.ceilingFan.Off()
}

func (this *CeilingFanOffCommand) Undo() {
	switch this.prevSpeed {
	case HIGH:
		this.ceilingFan.High()
	case MEDIUM:
		this.ceilingFan.Medium()
	case LOW:
		this.ceilingFan.Low()
	case OFF:
		this.ceilingFan.Off()

	}
}

//2.开吊扇
type CeilingFanOnCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOnCommand(ceilingFan *CeilingFan) *CeilingFanOnCommand {
	this := new(CeilingFanOnCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanOnCommand) Execute() {
	this.ceilingFan.High()
}

func (this *CeilingFanOnCommand) Undo() {
	this.ceilingFan.Off()
}

//3.开灯
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

func (this *LightOnCommand) Undo() {
	this.light.Off()
}

//4.关灯
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

func (this *LightOffCommand) Undo() {
	this.light.On()
}
