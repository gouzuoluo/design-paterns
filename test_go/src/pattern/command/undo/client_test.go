package romote

/*
*  客户请求命令（请求调用者，和请求接收者之前解耦）
*/
import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var remoteControl *RemoteControl = NewRemoteControl()

	//创建所有的设备device
	var livingRoomLight *Light = NewLight("Living Room")      //客厅灯
	var kitchenLight *Light = NewLight("Kitchen")             //厨房灯
	var ceilingFan *CeilingFan = NewCeilingFan("Living Room") //客厅吊扇


	//创建所有的灯命令
	var livingRoomLightOn *LightOnCommand = NewLightOnCommand(livingRoomLight)
	var livingRoomLightOff *LightOffCommand = NewLightOffCommand(livingRoomLight)
	var kitchenLightOn *LightOnCommand = NewLightOnCommand(kitchenLight)
	var kitchenLightOff *LightOffCommand = NewLightOffCommand(kitchenLight)

	//创建吊扇的开与关命令
	var ceilingFanOn *CeilingFanOnCommand = NewCeilingFanOnCommand(ceilingFan)
	var ceilingFanOff *CeilingFanOffCommand = NewCeilingFanOffCommand(ceilingFan)

	//现在已经有了全部命令，可以将他们全部加载到遥插槽中
	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)


	//打印遥控器插槽，及其插槽所指向的命令
	fmt.Println(remoteControl)

	//测试每个插槽的开关按钮
	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)

	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)

	remoteControl.UndoButtonWasPushed()

}
