package DesignPattern.Composite;

import java.util.Iterator;

/**
 * Created by Administrator on 2018/5/7.
 */

//菜单组件
public abstract class MenuComponent {
    //组合方法
    public void add(MenuComponent menuComponent) {
        throw new UnsupportedOperationException();
    }

    public void remove(MenuComponent menuComponent) {
        throw new UnsupportedOperationException();
    }

    public MenuComponent getChild(int i) {
        throw new UnsupportedOperationException();
    }

    //菜单项方法，其中一些也可以被组合使用
    public String getName() {
        throw new UnsupportedOperationException();
    }

    public String getDescription() {
        throw new UnsupportedOperationException();
    }

    public double getPrice() {
        throw new UnsupportedOperationException();
    }

    public boolean isVegetarian() {
        throw new UnsupportedOperationException();
    }

    //这个操作同时被菜单和菜单项所实现，这里提供默认操作
    public void print() {
        throw new UnsupportedOperationException();
    }

    public abstract Iterator<MenuComponent> createIterator();
}
