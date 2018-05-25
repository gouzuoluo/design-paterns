package simple_romote

//只有一个命令插槽，它是一个Invoker
type SimpleRemoteControl struct {
	slot Command
}

func (this *SimpleRemoteControl) SetCommand(command Command) {
	this.slot = command
}

func (this *SimpleRemoteControl) ButtonWasPressed() {
	this.slot.Execute()
}
