#单例模式

有一些对象我们只需要一个，比如：数据库连接，线程池（threadpool）、缓存（cache）、对话框、处理
偏好设置和注册表（registry）的对象、日志对象，充当打印机、显卡等设备的驱动程序的对象。事实上，
这类对象只能有一个实例，如果制造出多个实例，就会导致许多问题产生。

>单件模式：确保一个类只有一个实例，并提供一个全局访问点。

单例模式有2种实现形式：在静态初始化中就创建单件和延迟实例化（单件比较消耗资源时建议用此方法）

使用延迟实例化（懒惰模式）的单例模式，使用“双重检查加锁”保证线程安全

##要点

* 单件模式确保程序中一个类最多只有一个实例。
* 单件模式也提供访问这个实例的全局点。
* 在Java中实现单件模式需要私有的构造器，一个静态方法和一个静态变量。
* 确定在性能和资源上的限制，然后小心地选择适当的方案来实现单件，以解决多线程的问题。
* Java中如果有多个类加载器，加载同一个单件类，就会产生多个单件并存的怪异现象。此时可以自行指定类
加载器，并指定同一个类加载器。