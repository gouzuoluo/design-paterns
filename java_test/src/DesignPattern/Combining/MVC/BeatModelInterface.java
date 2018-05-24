package DesignPattern.Combining.MVC;

public interface BeatModelInterface {
   //这些方法是让控制器调用的，控制器根据用户的操作而对模型做出适当的处理
    void initialize();
    void on();
    void off();
    void setBPM(int bpm);
    int getBPM();

    //这些方法允许视图和控制器取得状态，并变成观察者
    void registerObserver(BeatObserver o);
    void removeObserver(BeatObserver o);
    void registerObserver(BPMOberver o);
    void removeObserver(BPMOberver o);
}
