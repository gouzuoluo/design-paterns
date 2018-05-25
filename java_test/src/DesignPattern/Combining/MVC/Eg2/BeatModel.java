package DesignPattern.Combining.MVC.Eg2;

import javax.sound.midi.*;
import java.util.*;

//节拍模块实例
public class BeatModel implements BeatModelInterface, MetaEventListener {
    Sequencer sequencer;//定序器(产生节拍)
    ArrayList<BeatObserver> beatObservers = new ArrayList<BeatObserver>();//BeatObserver类型的观察者
    ArrayList<BPMObserver> bpmObservers = new ArrayList<BPMObserver>();//BPMObserver类型的观察者
    int bpm = 90;
    Sequence sequence;
    Track track;//音轨

    @Override
    //初始化：设置定序器和节拍音轨
    public void initialize() {
        setUpMidi();
        buildTrackAndStart();
    }

    @Override
    //开：开始定序器，并设置默认的BPM
    public void on() {
        System.out.println("Starting the sequencer");
        sequencer.start();
        setBPM(90);
    }

    @Override
    //关：停止定序器
    public void off() {
        setBPM(0);
        sequencer.stop();
    }

    @Override
    //设置BPM：设置BPM实例变量，通知BPM观察者
    public void setBPM(int bpm) {
        this.bpm = bpm;
        sequencer.setTempoInBPM(getBPM());
        notifyBPMObservers();
    }

    @Override
    //获取BPM值
    public int getBPM() {

        return bpm;
    }

    @Override
    //注册BPMObserver
    public void registerObserver(BPMObserver o) {

        bpmObservers.add(o);
    }

    @Override
    //移除BPMObserver
    public void removeObserver(BPMObserver o) {
        int i = bpmObservers.indexOf(o);
        if (i >= 0) {
            bpmObservers.remove(i);
        }
    }

    public void notifyBPMObservers() {
        for (int i = 0; i < bpmObservers.size(); i++) {
            BPMObserver observer = (BPMObserver) bpmObservers.get(i);
            observer.updateBPM();
        }
    }

    @Override
    //注册BeatObserver
    public void registerObserver(BeatObserver o) {

        beatObservers.add(o);
    }

    @Override
    //移除BeatObserver
    public void removeObserver(BeatObserver o) {
        int i = beatObservers.indexOf(o);
        if (i >= 0) {
            beatObservers.remove(i);
        }
    }

    public void notifyBeatObservers() {
        for (int i = 0; i < beatObservers.size(); i++) {
            BeatObserver observer = (BeatObserver) beatObservers.get(i);
            observer.updateBeat();
        }
    }


    //设定定序器
    public void setUpMidi() {
        try {
            sequencer = MidiSystem.getSequencer();
            sequencer.open();
            sequencer.addMetaEventListener(this);

            sequence = new Sequence(Sequence.PPQ, 4);
            track = sequence.createTrack();
            sequencer.setTempoInBPM(getBPM());
            sequencer.setLoopCount(Sequencer.LOOP_CONTINUOUSLY);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    //创建音轨并开始
    public void buildTrackAndStart() {
        int[] trackList = {35, 0, 46, 0};

        sequence.deleteTrack(null);
        track = sequence.createTrack();

        makeTracks(trackList);
        track.add(makeEvent(192, 9, 1, 0, 4));
        try {
            sequencer.setSequence(sequence);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    //制作音轨
    public void makeTracks(int[] list) {

        for (int i = 0; i < list.length; i++) {
            int key = list[i];
            if (key != 0) {
                track.add(makeEvent(144, 9, key, 100, i));
                track.add(makeEvent(128, 9, key, 100, i + 1));
            }
        }
    }

    //制作事件
    public MidiEvent makeEvent(int comd, int chan, int one, int two, int tick) {
        MidiEvent event = null;
        try {
            ShortMessage a = new ShortMessage();
            a.setMessage(comd, chan, one, two);
            event = new MidiEvent(a, tick);

        } catch (Exception e) {
            e.printStackTrace();
        }
        return event;
    }

    @Override
    //节拍变化回调函数(TODO:这里有点问题，目前没有自动回调)
    public void meta(MetaMessage message) {
        if (message.getType() == 47) {
            beatEvent();
            sequencer.start();
            setBPM(getBPM());
        }
    }

    //打拍子事件
    void beatEvent() {

        notifyBeatObservers();
    }

}
