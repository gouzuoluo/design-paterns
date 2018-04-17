package DesignPattern.Command.SimpleRemote;

/**
 * Created by Administrator on 2018/4/16.
 */

//车库门打开命令
public class GarageDoorOpenCommand implements Command {
    GarageDoor garageDoor;

    public GarageDoorOpenCommand(GarageDoor garageDoor) {
        this.garageDoor = garageDoor;
    }

    public void execute() {
        garageDoor.up();
    }
}