package DesignPattern.Factory.AbstractFactory;

/**
 * Created by Administrator on 2018/4/8.
 */
public abstract class Pizza {
    String name;//名字
    Dough dough;//生面团
    Sauce sauce;//酱油
    Veggies veggies[];//蔬菜
    Cheese cheese;//奶酪
    Pepperoni pepperoni;//意大利香肠
    Clams clam;//蛤蜊

    abstract void prepare();//将该方法设置成抽象的，在这个方法中我们需要收集披萨所需的原料，而这写原料都来自原料工厂

    void bake() {
        System.out.println("Bake for 25 minutes at 350");
    }

    void cut() {
        System.out.println("Cutting the pizza into diagonal slices");
    }

    void box() {
        System.out.println("Place pizza in official PizzaStore box");
    }

    void setName(String name) {
        this.name = name;
    }

    String getName() {
        return name;
    }

    public String toString() {
        StringBuffer result = new StringBuffer();
        result.append("---- " + name + " ----\n");
        if (dough != null) {
            result.append(dough);
            result.append("\n");
        }
        if (sauce != null) {
            result.append(sauce);
            result.append("\n");
        }
        if (cheese != null) {
            result.append(cheese);
            result.append("\n");
        }
        if (veggies != null) {
            for (int i = 0; i < veggies.length; i++) {
                result.append(veggies[i]);
                if (i < veggies.length - 1) {
                    result.append(", ");
                }
            }
            result.append("\n");
        }
        if (clam != null) {
            result.append(clam);
            result.append("\n");
        }
        if (pepperoni != null) {
            result.append(pepperoni);
            result.append("\n");
        }
        return result.toString();
    }
}
