package eg1

import "sync"

//使用首字母小写，对包外不可见
type singleton struct {
}

//私有的创建函数，对包外不可见
func newSingleton() *singleton {
	this := new(singleton)
	return this
}

func (this *singleton) String() string {
	return "I am singleton"
}

var (
	singletonInstance *singleton = nil
	mu                sync.Mutex
)

/*----------------------------------------------------------*/

//提供全局访问点
func GetSingletonInstance() *singleton {
	//使用“双重判断加锁”来支持多线程
	if singletonInstance == nil {
		mu.Lock()
		if singletonInstance == nil {
			singletonInstance = newSingleton()
		}
		mu.Unlock()
	}
	return singletonInstance
}
