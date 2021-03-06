#观察者模式

出版者 + 订阅者 = 观察者模式

主题  + 观察者 = 观察者模式

* 主题对象管理某些数据，当主题内的数据改变，就会通知观察者，新的数据会以某种形式送到观察者手上。
* 观察者已经订阅（注册）主题以便在主题数据改变时能够收到更新。

##定义

> 观察者模式：定义了对象之间的一对多个依赖，这样一来，当一个对象改变状态时，它的所有依赖者都会收到
通知并自动更新。

主题和观察者定义了一对多的关系。观察者依赖于此主题，只要主题状态一有变化，观察者就会被通知。根据
通知的风格，观察者可能因此新值而更新。

主题是真正拥有数据的人，观察者是主题的依赖者，在数据变化时更新，这样比起让许多对象控制同一份数据
来，可以得到更干净的OO设计。

##松耦合的威力

当两个对象之间松耦合，它们依然可以交互，但是不太清楚彼此的细节。观察者模式提供了一种对象设计，让
主题和观察者之间松耦合。

* 关于观察者的一切，主题只知道观察者实现了某个接口（Observer接口）。主题不需要知道观察者的具体类
是谁、做了些什么或其他任何细节。
* 任何时候我们都可以增加新的观察者。同样，也可以在任何时候删除某些观察者。
* 在新类型的观察者出现时，主题的代码不需要更改。
* 我们可以独立的复用主题或观察者。如果我们在其他地方需要使用主题或观察者，可以轻易地复用，因为二者
并非紧耦合。
* 改变主题或者观察者其中一方，并不影响另一方。因为两者是松耦合，所以只要他们之间的接口仍被遵守，我
们就可以自由地改变他们。

##设计原则

* 设计原则4：为了交互对象之间的松耦合设计而努力。


##要点

* 主题（也就是可观察者）用一个共同的接口来更新观察者。
* 观察者和可观察者之间用松耦合方式结合，可观察者不知道观察者的细节，只知道观察者实现了观察者接口。
* 使用此模式时，你可从被观察者处推（push）或拉（pull）数据（然而，推的方式被认为更“正确”）。
* 有多个观察者时，不可以依赖特定的通知次序。