#Model-View-Controller(MVC)

控制器位于视图和模型之前。它将用户的输入，转给模型做动作。

##模型-视图-控制器

* 视图：用来呈现模型。视图通常直接从模型中取得它需要显示的状态与数据。

* 控制器：取得用户的输入并解读其对模型的意思。

* 模型：模型持有所有的数据、状态和程序逻辑。模型没有注意到视图和控制器，虽然它
提供了操纵和检索状态的接口，并发送状态改变通知给观察者。

控制器存在的主要原因：控制器把控制逻辑从视图中分离，让模型和视图之间解耦。通过保持控制器和视图之间的
松耦合，设计就更有弹性而且容易扩展，足以容纳以后的改变。

##从设计模式看

（策略）视图和控制器实现了经典的策略模式：视图是一个对象，可以被调整使用不同的策略，
而控制器提供了策略。视图只关心系统中可视的部分，对于任何界面行为，都委托给控制器处理。
使用策略模式也可以让视图和模型之间的关心解耦，因为控制器负责和模型交互来传递用户的
请求。对于工作是怎么完成的，视图毫不知情。对视图来说，控制器是策略，也就是知道如
何处理用户动作的对象。想换一种行为，换掉控制器就可以了。

（观察者）模型实现了观察者模式，当状态改变时，相关对象将持续更新。使用观察者模式，
可以让模型完全独立于视图和控制器。同一个模型可以使用不同的视图，甚至可以同时使用
多个视图。

（组合）视图内部使用组合模式来管理窗口、按钮以及其他显示组件。每个显示组件如果不是
组合节点（如窗口），就是叶节点（如按钮）。当控制器告诉视图更新时，只需要告诉视图
最顶层的组件即可，组合会处理其余的事情。
