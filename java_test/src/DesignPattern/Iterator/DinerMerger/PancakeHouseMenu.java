package DesignPattern.Iterator.DinerMerger;

import java.util.ArrayList;

/**
 * Created by Administrator on 2018/5/3.
 */

//煎饼屋菜单
public class PancakeHouseMenu implements Menu {
    ArrayList<MenuItem> menuItems;

    public PancakeHouseMenu() {
        menuItems = new ArrayList<MenuItem>();
        addItem("K&B's Pancake Breakfast",
                "Pancakes with scrambled eggs, and toast",
                true,
                2.99);

        addItem("Regular Pancake Breakfast",
                "Pancakes with fried eggs, sausage",
                false,
                2.99);

        addItem("Blueberry Pancakes",
                "Pancakes made with fresh blueberries",
                true,
                3.49);

        addItem("Waffles",
                "Waffles, with your choice of blueberries or strawberries",
                true,
                3.59);
    }

    //添加菜单项
    public void addItem(String name,
                        String description,
                        boolean vegetarian,
                        double price) {
        MenuItem menuItem = new MenuItem(name, description, vegetarian, price);
        menuItems.add(menuItem);
    }

    //获取菜单列表
    public ArrayList<MenuItem> getMenuItems() {
        return menuItems;
    }

    //实现Menu接口的createIterator方法
    public Iterator createIterator(){
        return new PancakeHouseMenuIterator(menuItems);
    }

    public String toString() {
        return "Objective Pancake House Menu";
    }

    // other menu methods here
}
