package _1_simple_romote

/*
*  测试客户请求命令（请求调用者，和请求接收者之前解耦）
*/
import "testing"

func TestAll(t *testing.T) {

	//构建开灯命令
	var lightOn *LightOnCommand = NewLightOnCommand(new(Light))

	//构建打开车库门命令
	var doorOpen *GarageDoorOpenCommand = NewGarageDoorOpenCommand(new(GarageDoor))

	//简单遥控
	var remote *SimpleRemoteControl = new(SimpleRemoteControl)

	//开灯
	remote.SetCommand(lightOn)
	remote.ButtonWasPressed()

	//开车库门
	remote.SetCommand(doorOpen)
	remote.ButtonWasPressed()
}
