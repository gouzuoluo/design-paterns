package DesignPattern.Adapter;

/**
 * Created by Administrator on 2018/4/18.
 */
public class DuckTestDrive {
    public static void main(String[] args){
        MallardDuck duck = new MallardDuck();//创建一只鸭子

        WildTurkey turkey = new WildTurkey();//创建一只火鸡

        Duck turkeyAdapter = new TurkeyAdapter(turkey);//将火鸡包装进一个火鸡适配器中，使它看起来像一只鸭子

        //测试这只火鸡，让它各个叫，让它飞行
        System.out.println("The Turkey says...");
        turkey.gobble();
        turkey.fly();

        //使用testDuck测试鸭子，需要传入一个鸭子对象
        System.out.println("\nThe Duck says...");
        testDuck(duck);


        //测试传入假装是鸭子的火鸡
        System.out.println("\nThe TurkeyAdapter says...");
        testDuck(turkeyAdapter);
    }


    static void testDuck(Duck duck){
        duck.quack();
        duck.fly();
    }

}
