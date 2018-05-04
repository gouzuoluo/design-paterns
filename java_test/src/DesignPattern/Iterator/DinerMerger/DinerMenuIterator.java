package DesignPattern.Iterator.DinerMerger;

/**
 * Created by Administrator on 2018/5/3.
 */

//餐厅菜单的迭代器
public class DinerMenuIterator implements Iterator {
    MenuItem[] items;
    int position = 0;

    public DinerMenuIterator(MenuItem[] items) {
        this.items = items;
    }

    //实现Iterator接口的next方法
    public MenuItem next() {
        MenuItem menuItem = items[position];
        position = position + 1;
        return menuItem;
    }

    //实现Iterator接口的hasNext方法
    public boolean hasNext() {
        if (position >= items.length || items[position] == null) {
            return false;
        } else {
            return true;
        }
    }
}
