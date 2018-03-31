package DesignPattern.Strategy;

/**
 * Created by Administrator on 2018/3/26.
 */
public class FlyRocketPowered implements FlyBehavior {
    public void fly(){
        System.out.println("I'm flying with a rocket!");
    }
}
