package DesignPattern.State.eg2;

/**
 * Created by Administrator on 2018/5/10.
 */
public interface State {

    //动作一：投入25美分
    public void insertQuarter();

    //动作二：退回25美分
    public void ejectQuarter();

    //动作三：转动曲柄
    public void turnCrank();

    //动作四：发放糖果()
    public void dispense();

    //动作五：重装糖果
    public void refill();
}
