package DesignPattern.Strategy;

/**
 * Created by Administrator on 2018/3/26.
 */
public class ModelDuck extends Duck {
    public ModelDuck(){
        flyBehavior = new FlyNoWay();
        quackBehavior = new Quack();
    }

    public void display(){
        System.out.println("I'm a model duck");
    }
}
