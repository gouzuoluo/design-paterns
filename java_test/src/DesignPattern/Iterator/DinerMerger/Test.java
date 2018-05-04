package DesignPattern.Iterator.DinerMerger;

import java.util.ArrayList;

/**
 * Created by Administrator on 2018/5/3.
 */
public class Test {
    public static void main(String args[]) {
        //创建各自的菜单
        PancakeHouseMenu pancakeHouseMenu = new PancakeHouseMenu();
        DinerMenu dinerMenu = new DinerMenu();

        // Without iterators
        //printMenu();

        //创建女服务员
        Waitress waitress = new Waitress(pancakeHouseMenu, dinerMenu);
        waitress.printMenu();
    }

    /*
     * Without the Waitress, we need the code below...
	 */
    public static void printMenu() {
        PancakeHouseMenu pancakeHouseMenu = new PancakeHouseMenu();
        DinerMenu dinerMenu = new DinerMenu();

        ArrayList<MenuItem> breakfastItems = pancakeHouseMenu.getMenuItems();
        MenuItem[] lunchItems = dinerMenu.getMenuItems();

        // Hiding implementation
        System.out.println("USING FOR EACH");
        for (MenuItem menuItem : breakfastItems) {
            System.out.print(menuItem.getName());
            System.out.println("\t\t" + menuItem.getPrice());
            System.out.println("\t" + menuItem.getDescription());
        }
        for (MenuItem menuItem : lunchItems) {
            System.out.print(menuItem.getName());
            System.out.println("\t\t" + menuItem.getPrice());
            System.out.println("\t" + menuItem.getDescription());
        }

        // Exposing implementation
        System.out.println("USING FOR LOOPS");
        for (int i = 0; i < breakfastItems.size(); i++) {
            MenuItem menuItem = (MenuItem) breakfastItems.get(i);
            System.out.print(menuItem.getName());
            System.out.println("\t\t" + menuItem.getPrice());
            System.out.println("\t" + menuItem.getDescription());
        }

        for (int i = 0; i < lunchItems.length; i++) {
            MenuItem menuItem = lunchItems[i];
            System.out.print(menuItem.getName());
            System.out.println("\t\t" + menuItem.getPrice());
            System.out.println("\t" + menuItem.getDescription());
        }
    }
}
