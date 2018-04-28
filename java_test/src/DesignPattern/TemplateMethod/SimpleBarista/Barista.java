package DesignPattern.TemplateMethod.SimpleBarista;

/**
 * Created by Administrator on 2018/4/28.
 */

//咖啡调配师
public class Barista {
    public static void main(String[] args){
        Tea tea = new Tea();
        Coffee coffee = new Coffee();

        System.out.println("Making tea...");
        tea.prepareRecipe();

        System.out.println("Making coffee...");
        coffee.prepareRecipe();
    }
}
