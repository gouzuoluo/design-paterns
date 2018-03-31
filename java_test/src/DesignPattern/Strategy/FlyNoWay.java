package DesignPattern.Strategy;

/**
 * Created by Administrator on 2018/3/26.
 */
public class FlyNoWay implements FlyBehavior {
    public void fly(){
        System.out.println("I can't fly");
    }
}
