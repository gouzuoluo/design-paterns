package DesignPattern.Composite;

import java.util.Iterator;

/**
 * Created by Administrator on 2018/5/7.
 */

//菜单项
public class MenuItem extends MenuComponent {
    String name;
    String description;
    boolean vegetarian;
    double price;

    public MenuItem(String name,
                    String description,
                    boolean vegetarian,
                    double price) {
        this.name = name;
        this.description = description;
        this.vegetarian = vegetarian;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    public double getPrice() {
        return price;
    }

    public boolean isVegetarian() {
        return vegetarian;
    }

    public String toString() {
        return (name + ", $" + price + "\n   " + description);
    }

    //重写MenuComponent类的print方法
    public void print() {
        System.out.print("  " + getName());
        if (isVegetarian()) {
            System.out.print("(v)");
        }
        System.out.println(", " + getPrice());
        System.out.println("   --" + getDescription());
    }

    public Iterator<MenuComponent> createIterator() {
        return new NullIterator();
    }

}
