package romote

/*
*  制造一种新的命令，用来执行其他一堆的命令。。。而不是执行一个命令
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
	var stereo *Stereo = NewStereo("Living Room")             //立体影响

	//创建所有的灯命令
	var livingRoomLightOn *LightOnCommand = NewLightOnCommand(livingRoomLight)
	var livingRoomLightOff *LightOffCommand = NewLightOffCommand(livingRoomLight)
	var kitchenLightOn *LightOnCommand = NewLightOnCommand(kitchenLight)
	var kitchenLightOff *LightOffCommand = NewLightOffCommand(kitchenLight)

	//创建吊扇的开与关命令
	var ceilingFanOn *CeilingFanOnCommand = NewCeilingFanOnCommand(ceilingFan)
	var ceilingFanOff *CeilingFanOffCommand = NewCeilingFanOffCommand(ceilingFan)

	//创建立体音响的开与关命令
	var stereoOnWithCD *StereoOnWithCDCommand = NewStereoOnWithCDCommand(stereo)
	var stereoOff *StereoOffCommand = NewStereoOffCommand(stereo)

	//创建批量命令
	var partyOnMacro *MacroCommand = NewMacroCommand([]Command{
		livingRoomLightOn,
		kitchenLightOn,
		ceilingFanOn,
		stereoOnWithCD,
	})
	var partyOffMacro *MacroCommand = NewMacroCommand([]Command{
		livingRoomLightOff,
		kitchenLightOff,
		ceilingFanOff,
		stereoOff,
	})

	//加载命令到指定的遥插槽中
	remoteControl.SetCommand(0, partyOnMacro, partyOffMacro)

	//打印遥控器插槽，及其插槽所指向的命令
	fmt.Println(remoteControl)

	//测试按钮
	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
}
