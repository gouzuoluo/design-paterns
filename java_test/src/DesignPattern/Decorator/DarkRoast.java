package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class DarkRoast extends Beverage {
    public DarkRoast(){
        description = "Dark Roast Coffee";
    }

    public double cost(){
        return 1.0;
    }
}
