package DesignPattern.Adapter.Ducks;

/**
 * Created by Administrator on 2018/4/18.
 */

//首先需要实现需要转换成的类型接口，也就是你客户所期望看到的接口
public class TurkeyAdapter implements Duck {
    Turkey turkey;

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
