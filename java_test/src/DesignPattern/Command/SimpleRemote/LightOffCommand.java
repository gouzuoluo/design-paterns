package DesignPattern.Command.SimpleRemote;

/**
 * Created by Administrator on 2018/4/16.
 */

//关灯命令
public class LightOffCommand implements Command {
    Light light;

    public LightOffCommand(Light light) {
        this.light = light;
    }

    public void execute() {
        light.off();
    }
}
