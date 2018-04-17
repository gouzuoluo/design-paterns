package DesignPattern.Command.SimpleRemote;

/**
 * Created by Administrator on 2018/4/13.
 */
public class SimpleRemoteControl {
    Command slot;//有一个插槽持有命令，而这个命令控制着一个装置

    public SimpleRemoteControl() {

    }

    public void setCommand(Command command) {
        slot = command;
    }

    public void buttonWasPressed() {
        slot.execute();
    }


}
