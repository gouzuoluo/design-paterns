package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public abstract class Beverage {
    String description = "Unknown Beverage";

    public String getDescription(){
        return description;
    }

    public abstract double cost();
}
