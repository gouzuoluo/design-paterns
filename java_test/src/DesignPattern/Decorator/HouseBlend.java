package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class HouseBlend extends Beverage {
    public HouseBlend(){
        description = "House Blend Coffee";
    }

    public double cost(){
        return .89;
    }
}
