package eg4

//命令接口
type Command interface {
	Execute()
}

/*===================================================================================================================*/

/*
* 2.开始命令
*/
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (this *StartCommand) Execute() {
	this.mb.Start()
}

/*
* 2.重启命令
*/
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (this *RebootCommand) Execute() {
	this.mb.Reboot()
}
