package DesignPattern.Combining.Duck.Eg6Observer;

public interface QuackObservable {
    public void registerObserver(Observer observer);

    public void notifyObservers();
}
