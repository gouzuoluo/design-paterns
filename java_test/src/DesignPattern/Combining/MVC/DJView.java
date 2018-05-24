package DesignPattern.Combining.MVC;

import javax.swing.*;
import java.awt.event.ActionListener;

public class DJView implements ActionListener,BeatObserver,BPMObserver {
    BeatModelInterface model;
    ControllerInterface controllerInterface;
    JFrame viewFrame;
    JPanel viewPanel;
    BeatBar beatBar;
    JLabel bpmOutputLabel;
    JFrame controlFrame;
    JPanel controlPanel;
    JLabel bpmLabel;
    JTextField bpmTextField;
    JButton setBPMButton;
    JButton increaseBPMButton;
    JButton decreaseBPMButton;
    JMenuBar menuBar;
    JMenu menu;
    JMenuItem startMenuItem;
    JMenuItem stopMenuItem;
}
