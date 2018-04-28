package DesignPattern.TemplateMethod.SimpleBarista;

/**
 * Created by Administrator on 2018/4/28.
 */

//茶
public class Tea {
    //准备食谱
    void prepareRecipe() {
        boilWater();
        steepTeaBag();
        pourInCup();
        addLemon();
    }

    //烧开水
    public void boilWater() {
        System.out.println("Boiling water");
    }

    //泡茶包
    public void steepTeaBag() {
        System.out.println("Steeping the tea");
    }

    //添加柠檬
    public void addLemon() {
        System.out.println("Adding Lemon");
    }

    //倒进杯子
    public void pourInCup() {
        System.out.println("Pouring into cup");
    }
}
