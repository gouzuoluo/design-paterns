package DesignPattern.Strategy;

/**
 * Created by Administrator on 2018/3/26.
 */
public abstract class Duck {
    FlyBehavior flyBehavior;//将飞行行为委托给该接口
    QuackBehavior quackBehavior;//将呱呱叫的行为委托给该接口

    public Duck() {
    }

    public abstract void display();

    public void preformFly() {
        flyBehavior.fly();
    }

    public void performQuack() {
        quackBehavior.quack();
    }

    public void swim() {
        System.out.println("All ducks float,even decoys!");
    }

    public void setFlyBehavior(FlyBehavior fb) {
        flyBehavior = fb;
    }

    public void setQuackBehavior(QuackBehavior qb) {
        quackBehavior = qb;
    }
}
