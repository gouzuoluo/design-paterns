package DesignPattern.TemplateMethod.Barista;

/**
 * Created by Administrator on 2018/4/28.
 */
public class Tea extends CaffeineBeverage {
    public void brew() {
        System.out.println("Steeping the tea");
    }

    public void addCondiments() {
        System.out.println("Adding Lemon");
    }
}
