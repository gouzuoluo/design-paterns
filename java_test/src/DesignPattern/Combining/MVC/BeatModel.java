package DesignPattern.Combining.MVC;

import javax.sound.midi.Sequencer;
import java.util.ArrayList;

public class BeatModel implements BeatModelInterface, MetaEventListener {
    Sequencer sequencer;//定序器对象知道如何产生真实的节拍（你听到的拍子）
    ArrayList beatObservers = new ArrayList();//观察节拍
    ArrayList bpmObservers = new ArrayList();//观察BPM改变
    int bpm = 90;
    //其他实例变量

    //此方法为我们设置定序器和节拍
    @Override
    public void initialize() {
        setUpMidi();
        buildTrackAndStart();
    }

    @Override
    public void on() {
        sequencer.start();
        setBPM(90);
    }

    @Override
    public void off() {
        setBPM(0);
        sequencer.stop();
    }

    @Override
    public void setBPM(int bpm) {
        this.bpm = bpm;
        sequencer.setTempoInBPM(getBPM());
        notifyBPMObservers();
    }

    @Override
    public int getBPM() {
        return bpm;
    }

    void beatEvent(){
        notifyBeatObservers();
    }

    @Override
    public void registerObserver(BeatObserver o) {

    }

    @Override
    public void removeObserver(BeatObserver o) {

    }

    @Override
    public void registerObserver(BPMOberver o) {

    }

    @Override
    public void removeObserver(BPMOberver o) {

    }
}
