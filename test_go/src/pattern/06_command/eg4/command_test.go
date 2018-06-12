package eg4

import "testing"

func TestAll(t *testing.T) {
	//创建动作的接收者，主板
	mb := &MotherBoard{}

	//创建命令
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	//创建动作的请求者
	box := NewBox(startCommand, rebootCommand)
	box.PressButton1()
	box.PressButton2()
}
