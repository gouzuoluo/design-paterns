package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class Mocha extends CondimentDecorator {
    Beverage beverage;

    public Mocha(Beverage beverage) {
        this.beverage = beverage;
    }

    public String getDescription() {
        return beverage.getDescription() + ",Mocha";
    }

    public double cost() {
        return beverage.cost() + .20;
    }
}
