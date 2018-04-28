package DesignPattern.Adapter;

/**
 * Created by Administrator on 2018/4/18.
 */

//首先需要实现需要转换成的类型接口，也就是你客户所期望看到的接口

//Duck是目标接口
//适配器实现了目标接口，并持有被适配者的实例
public class TurkeyAdapter implements Duck {
    Turkey turkey;//被适配者接口

    public TurkeyAdapter(Turkey turkey) {
        this.turkey = turkey;//接着，需要取得要适配的对象的引用
    }

    //实现我们需要实现接口中的所有方法。
    public void quack() {
        turkey.gobble();
    }

    public void fly() {
        for (int i = 0; i < 5; i++) {
            turkey.fly();
        }
    }
}
