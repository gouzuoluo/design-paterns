package DesignPattern.Strategy;

/**
 * Created by Administrator on 2018/3/26.
 */
public class MuteQuack implements QuackBehavior{
    public void quack(){
        System.out.println("<< Silence >>");
    }
}
