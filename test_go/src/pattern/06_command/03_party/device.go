package romote

import "fmt"

/*
*  动作的执行者，它知道如何进行必要的工作。任何类都可以当动作的执行者
*/

//1.电视
type TV struct {
	location string
	channel  int8
}

func NewTV(location string) *TV {
	this := new(TV)
	this.location = location
	return this
}

func (this *TV) On() {
	fmt.Println("TV is on")
}

func (this *TV) Off() {
	fmt.Println("TV is off")
}

func (this *TV) SetInputChannel() {
	this.channel = 3
	fmt.Println("Channel is set for VCR")
}

//2.立体音响
type Stereo struct {
	location string
}

func NewStereo(location string) *Stereo {
	this := new(Stereo)
	this.location = location
	return this
}

func (this *Stereo) On() {
	fmt.Println(this.location, " stereo is on")
}

func (this *Stereo) Off() {
	fmt.Println(this.location, " stereo is off")
}

func (this *Stereo) SetCD() {
	fmt.Println(this.location, " stereo is set for CD input")
}

func (this *Stereo) SetDVD() {
	fmt.Println(this.location, " stereo is set for DVD input")
}

func (this *Stereo) SetVolume(volume int) {
	fmt.Println(this.location, " Stereo volume set to ", volume)
}

//3.电灯
type Light struct {
	location string
}

func NewLight(location string) *Light {
	this := new(Light)
	this.location = location
	return this
}

func (this *Light) On() {
	fmt.Println(this.location, " light is on")
}

func (this *Light) Off() {
	fmt.Println(this.location, " light is off")
}

//4.热水浴缸
type HotTub struct {
	on          bool
	temperature int
}

func (this *HotTub) On() {
	this.on = true
}

func (this *HotTub) Off() {
	this.on = false
}

func (this *HotTub) BubblesOn() {
	if this.on {
		fmt.Println("Hot tub is bubbling!")
	}
}

func (this *HotTub) BubblesOff() {
	if this.on {
		fmt.Println("Hot tub is not bubbling")
	}
}

func (this *HotTub) JetsOn() {
	if this.on {
		fmt.Println("Hot tub jets are on")
	}
}

func (this *HotTub) SetTemperature(temperature int) {
	this.temperature = temperature
}

func (this *HotTub) Heat() {
	this.temperature = 105
	fmt.Println("Hot tub is heating to a steaming 105 degrees")
}

func (this *HotTub) Cool() {
	this.temperature = 98
	fmt.Println("Hot tub is cooling to 98 degrees")
}

//5.车库门
type GarageDoor struct {
	location string
}

func NewGarageDoor(location string) *GarageDoor {
	this := new(GarageDoor)
	this.location = location
	return this
}

func (this *GarageDoor) Up() {
	fmt.Println(this.location + " garage Door is Up")
}

func (this *GarageDoor) Down() {
	fmt.Println(this.location + " garage Door is Down")
}

func (this *GarageDoor) Stop() {
	fmt.Println(this.location + " garage Door is Stopped")
}

func (this *GarageDoor) LightOn() {
	fmt.Println(this.location + " garage light is on")
}

func (this *GarageDoor) LightOff() {
	fmt.Println(this.location + " garage light is off")
}

//6.吊扇
const (
	HIGH   int = 3
	MEDIUM int = 2
	LOW    int = 1
	OFF    int = 0
)

type CeilingFan struct {
	location string
	speed    int
}

func NewCeilingFan(location string) *CeilingFan {
	this := new(CeilingFan)
	this.location = location
	this.speed = OFF
	return this
}

func (this *CeilingFan) High() {
	this.speed = HIGH
	fmt.Println(this.location + " ceiling fan is on high")
}

func (this *CeilingFan) Medium() {
	this.speed = MEDIUM
	fmt.Println(this.location + " ceiling fan is on medium")
}

func (this *CeilingFan) Low() {
	this.speed = LOW
	fmt.Println(this.location + " ceiling fan is on low")
}

func (this *CeilingFan) Off() {
	this.speed = OFF
	fmt.Println(this.location + " ceiling fan is off")
}

func (this *CeilingFan) GetSpeed() int {
	return this.speed
}
