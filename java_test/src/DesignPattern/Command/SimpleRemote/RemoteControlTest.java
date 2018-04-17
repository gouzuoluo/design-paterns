package DesignPattern.Command.SimpleRemote;

/**
 * Created by Administrator on 2018/4/13.
 */
public class RemoteControlTest {
    public static void main(String[] args) {

        //构建打开灯命令
        Light light = new Light();
        LightOnCommand lightOn = new LightOnCommand(light);

        //构建车库门打开命令
        GarageDoor garageDoor = new GarageDoor();
        GarageDoorOpenCommand garageDoorOpen = new GarageDoorOpenCommand(garageDoor);

        /*-----------------------------------------------------------------------*/
        //简单遥控
        SimpleRemoteControl remote = new SimpleRemoteControl();

        //开灯
        remote.setCommand(lightOn);
        remote.buttonWasPressed();

        //开车库门
        remote.setCommand(garageDoorOpen);
        remote.buttonWasPressed();
    }
}
