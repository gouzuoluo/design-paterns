package _1_simple_romote

import "fmt"

/*
*  动作的执行者，它知道如何进行必要的工作。任何类都可以当动作的执行者
*/

/*
* 1.电灯设备
 */
type Light struct {
}

func (this *Light) On() {
	fmt.Println("Light is on")
}

func (this *Light) Off() {
	fmt.Println("Light is off")
}

/*
* 2.车库门设备
 */
type GarageDoor struct {
}

func (this *GarageDoor) Up() {
	fmt.Println("Garage Door is Open")
}

func (this *GarageDoor) Down() {
	fmt.Println("Garage Door is Closed")
}

func (this *GarageDoor) Stop() {
	fmt.Println("Garage Door is Stopped")
}

func (this *GarageDoor) LightOn() {
	fmt.Println("Garage light is on")
}

func (this *GarageDoor) LightOff() {
	fmt.Println("Garage light is off")
}
