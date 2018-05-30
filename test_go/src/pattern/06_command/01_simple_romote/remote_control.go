package _1_simple_romote

//只有一个命令插槽的遥控器，它是一个Invoker
type SimpleRemoteControl struct {
	slot Command
}

func (this *SimpleRemoteControl) SetCommand(command Command) {
	this.slot = command
}

func (this *SimpleRemoteControl) ButtonWasPressed() {
	this.slot.Execute()
}
