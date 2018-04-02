package DesignPattern.Decorator;

/**
 * Created by Administrator on 2018/4/2.
 */
public class StartbuzzCoffee {
    public static void main(String args[]) {
        Beverage beverage = new Espresso();
        System.out.println(beverage.getDescription() + " $" + beverage.cost());


        Beverage beverage1 = new DarkRoast();
        beverage1 = new Mocha(beverage);
        beverage1 = new Mocha(beverage1);
        beverage1 = new Whip(beverage1);
        System.out.println(beverage1.getDescription() + " $" + beverage1.cost());
    }
}
