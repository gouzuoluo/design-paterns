package _2_romote

/*
*  客户请求命令（请求调用者，和请求接收者之前解耦）
*/
import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	//遥控器
	var remoteControl *RemoteControl = NewRemoteControl()

	//创建所有的设备device
	var light *Light = NewLight("Living Room")                //灯
	var ceilingFan *CeilingFan = NewCeilingFan("Living Room") //吊扇


	//创建所有的灯命令
	var lightOn *LightOnCommand = NewLightOnCommand(light)
	var lightOff *LightOffCommand = NewLightOffCommand(light)


	//创建吊扇相关关命令
	var ceilingFanHigh *CeilingFanHighCommand = NewCeilingFanHighCommand(ceilingFan)
	var ceilingFanMedium *CeilingFanMediumCommand = NewCeilingFanMediumCommand(ceilingFan)
	var ceilingFanLow *CeilingFanLowCommand = NewCeilingFanLowCommand(ceilingFan)
	var ceilingFanOff *CeilingFanOffCommand = NewCeilingFanOffCommand(ceilingFan)


	//现在已经有了全部命令，可以将他们全部加载到遥插槽中
	remoteControl.SetCommand(0, lightOn, lightOff)
	remoteControl.SetCommand(1, ceilingFanHigh, ceilingFanMedium)
	remoteControl.SetCommand(2, ceilingFanLow, ceilingFanOff)

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
	remoteControl.UndoButtonWasPushed()
	remoteControl.OffButtonWasPushed(2)
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
}
