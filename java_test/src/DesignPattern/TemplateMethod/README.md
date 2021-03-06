# 模板方法模式
模板方法定义了一个算法的步骤，并允许子类为一个或多个步骤提供实现。

## 定义模板方法

模板方法模式：在一个方法中定义一个算法的骨架，而将一些步骤延长到子类中。模板方法使得子类可以在
不改变算法结构的情况下，重新定义算法中的某些步骤。

* 这个模式是用来创建一个算法的模板。
* 模板就是一个方法，更具体来说，这个方法定义成一组步骤，其中任何步骤都可以是抽象的，由
子类负责实现。这可以确保算法的结构保持不变，同时由子类提供部分实现。

## 钩子(hook)
钩子是一种被声明在抽象类中的方法，但只有空的或者默认的实现。

钩子的存在，可以让子类有能力对算法的不同点进行挂钩，要不要挂钩，由子类自习决定。
默认不做是的方法，我们称这种方法为“hook”,子类可以视情况决定要不要覆盖他们。

钩子的用途：
* 钩子可以让子类实现算法中的可选部分。
* 是让子类能够有机会对模板方法中某些即将发生的（或刚刚发生的）步骤做出反应。


当你的子类“必须”提供算法中某个方法或步骤的实现时，就使用抽象方法。如果算法的这个部分是
可选的，就用钩子。


## 好莱坞原则

别调用（打电话给）我们，我们会调用（打电话给）你。

在好莱坞原则之下，我们允许低层组件将自己挂钩到系统上，但是高层组件会决定什么时候和怎样使用
这些低层组件。


**模板方法模式是一种很常见的模式，因为对创建框架来说，这个模式简直棒极了。由框架控制如何和做事情，而由你（使用框架的人）指定框架算法中每个步骤的细节**

# 用模板方法排序
