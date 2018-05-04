package DesignPattern.Iterator.DinerMerger;

import java.util.ArrayList;

/**
 * Created by Administrator on 2018/5/3.
 */

//煎饼屋菜单的迭代器
public class PancakeHouseMenuIterator implements Iterator {
    ArrayList<MenuItem> items;
    int position = 0;

    public PancakeHouseMenuIterator(ArrayList<MenuItem> items) {
        this.items = items;
    }

    //实现Iterator接口的next方法
    public MenuItem next() {
        MenuItem item = items.get(position);
        position = position + 1;
        return item;
    }

    //实现Iterator接口的hasNext方法
    public boolean hasNext() {
        if (position >= items.size()) {
            return false;
        } else {
            return true;
        }
    }
}
