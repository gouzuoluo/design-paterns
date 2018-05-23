package DesignPattern.Combining.Duck.Eg6Observer;

//需要确定所有的Quackable都实现QuackObservable接口。所有干脆让其扩展于QuackObservable接口
public interface Quackable extends QuackObservable {
    public void quack();
}
