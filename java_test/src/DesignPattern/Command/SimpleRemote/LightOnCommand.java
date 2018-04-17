package DesignPattern.Command.SimpleRemote;

/**
 * Created by Administrator on 2018/4/13.
 */

//开灯命令
public class LightOnCommand implements Command {
    Light light;

    //构造器被传入了某个电灯，以便让这个命令控制，然后记录在实例变量中
    public LightOnCommand(Light light) {
        this.light = light;
    }

    public void execute() {
        light.on();
    }
}
