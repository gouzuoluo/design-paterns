package DesignPattern.Observer;


/**
 * Created by Administrator on 2018/3/29.
 */
public interface Subject {
    public void registerObserver(Observer o);
    public void removeObserver(Observer o);
    public void notifyObservers();
}
