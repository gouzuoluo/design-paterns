package DesignPattern.Combining.MVC.Eg1;

//控制器接口
public interface ControllerInterface {
    void start();
    void stop();
    void increaseBPM();
    void decreaseBPM();
    void setBPM(int bpm);
}