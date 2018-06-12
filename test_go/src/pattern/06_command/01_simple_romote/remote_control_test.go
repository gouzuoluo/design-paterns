package _1_simple_romote

/*
*  测试客户请求命令（请求调用者，和请求接收者之前解耦）
*/
import "testing"

func TestAll(t *testing.T) {

	//简单遥控
	var remote *SimpleRemoteControl = new(SimpleRemoteControl)

	//1.开灯
	var lightOn *LightOnCommand = NewLightOnCommand(new(Light)) //构建开灯命令
	remote.SetCommand(lightOn)
	remote.ButtonWasPressed()

	//2.开车库门
	var doorOpen *GarageDoorOpenCommand = NewGarageDoorOpenCommand(new(GarageDoor)) //构建打开车库门命令
	remote.SetCommand(doorOpen)
	remote.ButtonWasPressed()
}
