package romote

import "fmt"

/*
*  命令接收者。这里是各种device
*/


//1.电灯
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

//2.吊扇
const (
	HIGH   int = 3
	MEDIUM int = 2
	LOW    int = 1
	OFF    int = 0
)

type CeilingFan struct {
	location string
	level    int
}

func NewCeilingFan(location string) *CeilingFan {
	this := new(CeilingFan)
	this.location = location
	return this
}

func (this *CeilingFan) High() {
	this.level = HIGH
	fmt.Println(this.location + " ceiling fan is on high")
}

func (this *CeilingFan) Medium() {
	this.level = MEDIUM
	fmt.Println(this.location + " ceiling fan is on medium")
}

func (this *CeilingFan) Low() {
	this.level = LOW
	fmt.Println(this.location + " ceiling fan is on low")
}

func (this *CeilingFan) Off() {
	this.level = OFF
	fmt.Println(this.location + " ceiling fan is off")
}

func (this *CeilingFan) GetSpeed() int {
	return this.level
}
