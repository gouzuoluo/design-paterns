package DesignPattern.Adapter.Ducks;

/**
 * Created by Administrator on 2018/4/18.
 */
public class WildTurkey implements Turkey {
    public void gobble() {
        System.out.println("Gobble gobble");
    }

    public void fly() {
        System.out.println("I`m flying a short distance");
    }
}
