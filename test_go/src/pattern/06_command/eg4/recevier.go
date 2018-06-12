package eg4

import "fmt"

//主机板
type MotherBoard struct {
}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}
