package DesignPattern.Combined.MVC.Eg2;

//节拍控制器（Controller）
public class BeatController implements ControllerInterface {
    //控制器是MVC夹心饼中间的奶油，所有它必须同时和模型以及视图接触，来当两者的黏着剂
    BeatModelInterface model;
    DJView view;

    //构造函数
    public BeatController(BeatModelInterface model) {
        this.model = model;
        view = new DJView(this, model);
        view.createView();
        view.createControls();
        view.disableStopMenuItem();
        view.enableStartMenuItem();
        model.initialize();
    }

    //当用户从用户界面菜单中选择“start”时，控制器调用模型的on方法，然后改变用户界面（将start菜单项disable，将stop菜单项enable）
    public void start() {
        model.on();
        view.disableStartMenuItem();
        view.enableStopMenuItem();
    }

    //当用户从菜单中选择“stop”时，控制器调用模型的off方法，然后改变用户界面
    public void stop() {
        model.off();
        view.disableStopMenuItem();
        view.enableStartMenuItem();
    }

    /*
     * PS:控制器等于是在帮视图做决定。视图只知道如何将菜单项变成开和等，但是它并不知道在何种情况下要enable/disable
     */

    //点击增加按钮，控制器就从模型取得当前的BPM，加1，然后设置一个新的BPM
    public void increaseBPM() {
        int bpm = model.getBPM();
        model.setBPM(bpm + 1);
    }

    //点击增加按钮，控制器就从模型取得当前的BPM，减1，然后设置一个新的BPM
    public void decreaseBPM() {
        int bpm = model.getBPM();
        model.setBPM(bpm - 1);
    }

    //设置任意BPM值
    public void setBPM(int bpm) {

        model.setBPM(bpm);
    }
}
