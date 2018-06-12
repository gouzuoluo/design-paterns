package romote

import (
	"bytes"
	"reflect"
	"fmt"
)

//遥控器（动作的请求者invoker）
type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
}

func NewRemoteControl() *RemoteControl {
	this := new(RemoteControl)
	this.onCommands = make([]Command, 7)
	this.offCommands = make([]Command, 7)
	var noCommand Command = new(NoCommand)
	for i := 0; i < 7; i++ {
		this.onCommands[i] = noCommand
		this.offCommands[i] = noCommand
	}
	return this
}

func (this *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	this.onCommands[slot] = onCommand
	this.offCommands[slot] = offCommand
}

func (this *RemoteControl) OnButtonWasPushed(slot int) {
	this.onCommands[slot].Execute()
}

func (this *RemoteControl) OffButtonWasPushed(slot int) {
	this.offCommands[slot].Execute()
}

func (this *RemoteControl) String() string {
	buf := bytes.NewBufferString("\n------ Remote Control -------\n")
	for i := 0; i < len(this.onCommands); i++ {
		buf.WriteString(fmt.Sprintf("[slot %d] ", i) + reflect.TypeOf(this.onCommands[i]).String() +
			"    " + reflect.TypeOf(this.offCommands[i]).String() + "\n")
	}
	return buf.String()
}
