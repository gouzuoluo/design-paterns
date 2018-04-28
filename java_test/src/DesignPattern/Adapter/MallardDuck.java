package DesignPattern.Adapter;

/**
 * Created by Administrator on 2018/4/18.
 */
public class MallardDuck implements Duck {
    public void quack() {
        System.out.println("Quack");
    }

    public void fly() {
        System.out.println("I`m flying");
    }
}
