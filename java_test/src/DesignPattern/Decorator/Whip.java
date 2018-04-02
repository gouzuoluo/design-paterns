package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class Whip extends Beverage {
    Beverage beverage;

    public Whip(Beverage beverage) {
        this.beverage = beverage;
    }

    public String getDescription() {
        return beverage.getDescription() + ",Whip";
    }

    public double cost() {
        return beverage.cost() + .8;
    }
}
