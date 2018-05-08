package DesignPattern.Composite;

import java.util.Iterator;

/**
 * Created by Administrator on 2018/5/7.
 */
public class NullIterator implements Iterator<MenuComponent> {
    public MenuComponent next() {
        return null;
    }

    public boolean hasNext() {
        return false;
    }
}
