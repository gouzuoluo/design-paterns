package DesignPattern.TemplateMethod.SimpleBarista;

/**
 * Created by Administrator on 2018/4/28.
 */
public class Coffee {
    //准备食谱
    void prepareRecipe() {
        boilWater();
        brewCoffeeGrinds();
        pourInCup();
        addSugarAndMilk();
    }

    //烧水
    public void boilWater() {
        System.out.println("Boiling water");
    }

    //冲咖啡粉
    public void brewCoffeeGrinds() {
        System.out.println("Dripping Coffee through filter");
    }

    //倒进杯子
    public void pourInCup() {
        System.out.println("Pouring into cup");
    }

    //添加糖和牛奶
    public void addSugarAndMilk() {
        System.out.println("Adding Sugar and Milk");
    }
}
