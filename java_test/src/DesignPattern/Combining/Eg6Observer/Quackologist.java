package DesignPattern.Combining.Eg6Observer;

//实现一个观察者，呱呱叫学家
public class Quackologist implements Observer {
    @Override
    public void update(QuackObservable duck) {
        //简单地打印出正在呱呱叫的Quackable的对象
        System.out.println("Quackologist: " + duck + " just quacked.");
    }

    public String toString() {
        return "Quackologist";
    }
}
