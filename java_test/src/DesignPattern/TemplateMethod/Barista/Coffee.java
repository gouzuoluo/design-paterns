package DesignPattern.TemplateMethod.Barista;

/**
 * Created by Administrator on 2018/4/28.
 */
public class Coffee extends CaffeineBeverage {
    public void brew() {
        System.out.println("Dripping Coffee through filter");
    }

    public void addCondiments() {
        System.out.println("Adding Sugar and Milk");
    }
}
