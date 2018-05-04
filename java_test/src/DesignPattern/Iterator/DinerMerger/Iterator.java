package DesignPattern.Iterator.DinerMerger;

/**
 * Created by Administrator on 2018/5/3.
 */

//迭代器接口
public interface Iterator {
    boolean hasNext();//还有下一个菜单项

    MenuItem next();//获取下一个菜单项
}
