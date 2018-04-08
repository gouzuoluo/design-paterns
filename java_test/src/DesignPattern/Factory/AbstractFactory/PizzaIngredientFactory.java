package DesignPattern.Factory.AbstractFactory;

/**
 * Created by Administrator on 2018/4/8.
 */

//披萨原料工厂接口
public interface PizzaIngredientFactory {
    public Dough createDough();//创建面团

    public Sauce createSauce();//创建酱油

    public Cheese createCheese();//创建奶酪

    public Veggies[] createVeggies();//创建蔬菜

    public Pepperoni createPepperoni();//创建意大利香肠

    public Clams createClam();//创建蛤蜊
}
