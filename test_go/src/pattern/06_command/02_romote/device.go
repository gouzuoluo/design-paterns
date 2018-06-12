package _2_romote

import "fmt"

/*
*  动作的执行者，它知道如何进行必要的工作。任何类都可以当动作的执行者
*/

/*
* 1.电灯
 */
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

/*
* 2.吊扇
 */
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
