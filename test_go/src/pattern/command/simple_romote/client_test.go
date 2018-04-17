package simple_romote

/*
*  测试客户请求命令（请求调用者，和请求接收者之前解耦）
*/
import "testing"

func TestAll(t *testing.T) {

	//这个Client负责创建一个ConcreteCommand,并设置接收者

	//构建开灯命令
	var light *Light = new(Light)
	var lightOn *LightOnCommand = NewLightOnCommand(light)

	//构建打开车库门命令
	var door *GarageDoor = new(GarageDoor)
	var doorOpen *GarageDoorOpenCommand = NewGarageDoorOpenCommand(door)

	//简单遥控
	var remote *SimpleRemoteControl = new(SimpleRemoteControl)

	//开灯
	remote.SetCommand(lightOn)
	remote.ButtonWasPressed()

	//开车库门
	remote.SetCommand(doorOpen)
	remote.ButtonWasPressed()
}
