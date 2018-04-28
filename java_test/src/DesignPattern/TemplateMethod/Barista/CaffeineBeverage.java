package DesignPattern.TemplateMethod.Barista;

/**
 * Created by Administrator on 2018/4/28.
 */

//咖啡因饮料
public abstract class CaffeineBeverage {
    //准备食谱
    final void prepareRecipe() {
        boilWater();
        brew();
        pourInCup();
        addCondiments();
    }

    //冲泡
    abstract void brew();

    //添加调料
    abstract void addCondiments();

    //烧水
    void boilWater() {
        System.out.println("Boiling water");
    }

    //水倒进杯子
    void pourInCup() {
        System.out.println("Pouring into cup");
    }
}
