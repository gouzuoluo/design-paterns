package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class Espresso extends Beverage {
    public Espresso() {
        description = "Espresso";
    }

    public double cost() {
        return 1.99;
    }
}
