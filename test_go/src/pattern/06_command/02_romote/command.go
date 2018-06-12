package _2_romote

//命令接口
type Command interface {
	Execute()

	Undo() //撤销
}

/*===================================================================================================================*/

/*
* 0.无具体操作命令
 */
type NoCommand struct {
}

func (this *NoCommand) Execute() {
}

func (this *NoCommand) Undo() {
}

/*
* 1.开灯
 */
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

/*
* 2.关灯
 */
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

/*
* 3.高速运行吊扇
 */
type CeilingFanHighCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int //增加局部状态以便追踪吊扇之前的速度
}

func NewCeilingFanHighCommand(ceilingFan *CeilingFan) *CeilingFanHighCommand {
	this := new(CeilingFanHighCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanHighCommand) Execute() {
	this.prevSpeed = this.ceilingFan.GetSpeed()
	this.ceilingFan.High()
}

func (this *CeilingFanHighCommand) Undo() {
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

/*
* 4.中速运行吊扇
 */
type CeilingFanMediumCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int //增加局部状态以便追踪吊扇之前的速度
}

func NewCeilingFanMediumCommand(ceilingFan *CeilingFan) *CeilingFanMediumCommand {
	this := new(CeilingFanMediumCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanMediumCommand) Execute() {
	this.prevSpeed = this.ceilingFan.GetSpeed()
	this.ceilingFan.Medium()
}

func (this *CeilingFanMediumCommand) Undo() {
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

/*
* 5.低速运行吊扇
 */
type CeilingFanLowCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int //增加局部状态以便追踪吊扇之前的速度
}

func NewCeilingFanLowCommand(ceilingFan *CeilingFan) *CeilingFanLowCommand {
	this := new(CeilingFanLowCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanLowCommand) Execute() {
	this.prevSpeed = this.ceilingFan.GetSpeed()
	this.ceilingFan.Low()
}

func (this *CeilingFanLowCommand) Undo() {
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

/*
* 6.关闭吊扇
 */
type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int //增加局部状态以便追踪吊扇之前的速度
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

