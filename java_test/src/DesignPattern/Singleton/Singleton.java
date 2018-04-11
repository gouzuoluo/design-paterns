package DesignPattern.Singleton;

/**
 * Created by Administrator on 2018/4/11.
 */
public class Singleton {
    private volatile static Singleton uniqueInstance;
    private Singleton(){}
    public static Singleton  getInstance(){
        if (uniqueInstance == null) {
            synchronized (Singleton.class){
                if (uniqueInstance == null) {
                    uniqueInstance = new Singleton();
                }
            }
        }
        return uniqueInstance;
    }
}
