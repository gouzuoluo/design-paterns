package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class Decaf extends Beverage {
    public Decaf(){
        description = "Decaf Coffee";
    }

    public double cost(){
        return .88;
    }
}
