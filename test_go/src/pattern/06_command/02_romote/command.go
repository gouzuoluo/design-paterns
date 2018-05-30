package _2_romote

//命令接口
type Command interface {
	Execute()
}

/*===================================================================================================================*/

/*
* 0.无具体操作命令
 */
type NoCommand struct {
}

func (this *NoCommand) Execute() {
	//什么也不干
}

/*
* 1.关闭吊扇
 */
type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) *CeilingFanOffCommand {
	this := new(CeilingFanOffCommand)
	this.ceilingFan = ceilingFan
	return this
}

func (this *CeilingFanOffCommand) Execute() {
	this.ceilingFan.Off()
}

/*
* 2.开吊扇
 */
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

/*
* 3.开启车库门
 */
type GarageDoorUpCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorUpCommand(garageDoor *GarageDoor) *GarageDoorUpCommand {
	this := new(GarageDoorUpCommand)
	this.garageDoor = garageDoor
	return this
}

func (this *GarageDoorUpCommand) Execute() {
	this.garageDoor.Up()
}

/*
* 4.放下车库门
 */
type GarageDoorDownCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorDownCommand(garageDoor *GarageDoor) *GarageDoorDownCommand {
	this := new(GarageDoorDownCommand)
	this.garageDoor = garageDoor
	return this
}

func (this *GarageDoorDownCommand) Execute() {
	this.garageDoor.Down()
}

/*
* 5.打开热水浴缸
 */
type HotTubOnCommand struct {
	hotTub *HotTub
}

func NewHotTubOnCommand(hotTub *HotTub) *HotTubOnCommand {
	this := new(HotTubOnCommand)
	this.hotTub = hotTub
	return this
}

func (this *HotTubOnCommand) Execute() {
	this.hotTub.On()
	this.hotTub.Heat()
	this.hotTub.BubblesOn()
}

/*
*6.关闭热水浴缸
 */
type HotTubOffCommand struct {
	hotTub *HotTub
}

func NewHotTubOffCommand(hotTub *HotTub) *HotTubOffCommand {
	this := new(HotTubOffCommand)
	this.hotTub = hotTub
	return this
}

func (this *HotTubOffCommand) Execute() {
	this.hotTub.Cool()
	this.hotTub.Off()
}

/*
* 7.开灯
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

/*
* 8.关灯
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

/*
* 9.关闭客厅灯
 */
type LivingRoomLightOffCommand struct {
	light *Light
}

func NewLivingRoomLightOffCommand(light *Light) *LivingRoomLightOffCommand {
	this := new(LivingRoomLightOffCommand)
	this.light = light
	return this
}

func (this *LivingRoomLightOffCommand) Execute() {
	this.light.Off()
}

/*
* 10.打开客厅灯
 */
type LivingRoomLightOnCommand struct {
	light *Light
}

func NewLivingRoomLightOnCommand(light *Light) *LivingRoomLightOnCommand {
	this := new(LivingRoomLightOnCommand)
	this.light = light
	return this
}

func (this *LivingRoomLightOnCommand) Execute() {
	this.light.On()
}

/*
* 11.关闭立体音响
 */
type StereoOffCommand struct {
	stereo *Stereo
}

func NewStereoOffCommand(stereo *Stereo) *StereoOffCommand {
	this := new(StereoOffCommand)
	this.stereo = stereo
	return this
}

func (this *StereoOffCommand) Execute() {
	this.stereo.Off()
}

/*
* 12.播放CD
 */
type StereoOnWithCDCommand struct {
	stereo *Stereo
}

func NewStereoOnWithCDCommand(stereo *Stereo) *StereoOnWithCDCommand {
	this := new(StereoOnWithCDCommand)
	this.stereo = stereo
	return this
}

func (this *StereoOnWithCDCommand) Execute() {
	this.stereo.On()
	this.stereo.SetCD()
	this.stereo.SetVolume(11)
}
