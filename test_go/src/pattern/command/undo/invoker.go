package romote

import (
	"bytes"
	"reflect"
	"fmt"
)

//遥控器，this is the invoker
type RemoteControl struct {
	onCommands  []Command //开插槽命令
	offCommands []Command //关插槽命令
	undoCommand Command   //前一个命令将被记录在此
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
	this.undoCommand = noCommand
	return this
}

func (this *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	this.onCommands[slot] = onCommand
	this.offCommands[slot] = offCommand
}

func (this *RemoteControl) OnButtonWasPushed(slot int) {
	this.onCommands[slot].Execute()
	this.undoCommand = this.onCommands[slot]
}

func (this *RemoteControl) OffButtonWasPushed(slot int) {
	this.offCommands[slot].Execute()
	this.undoCommand = this.offCommands[slot]
}

func (this *RemoteControl) UndoButtonWasPushed() {
	this.undoCommand.Undo()
}

func (this *RemoteControl) String() string {
	buf := bytes.NewBufferString("\n------ Remote Control -------\n")
	for i := 0; i < len(this.onCommands); i++ {
		buf.WriteString(fmt.Sprintf("[slot %d] ", i) + reflect.TypeOf(this.onCommands[i]).String() +
			"    " + reflect.TypeOf(this.offCommands[i]).String() + "\n")
	}
	return buf.String()
}
