package DesignPattern.proxy.eg1;

/**
 * Created by Administrator on 2018/5/10.
 */
public class Test {
    public static void main(String[] args) {

        String location = "aaa";
        int count = 1;

        GumballMachine gumballMachine = new GumballMachine(location, count);

        GumballMonitor monitor = new GumballMonitor(gumballMachine);

        monitor.report();

    }
}
