package eg1

type Subject interface {
	Do() string
}

/*====================================================================================================================*/

//1.真实类
type RealSubject struct {
}

func (*RealSubject) Do() string {
	return "real"
}

//2.代理类
type Proxy struct {
	realSubject *RealSubject
}

func NewProxy()*Proxy  {
	this := new(Proxy)
	this.realSubject = new(RealSubject)
	return this
}

func (this *Proxy) Do() string {
	var res string
	// 在调用真实对象之前的工作(检查缓存，判断权限，实例化真实对象等。。)
	res += "pre:"

	// 调用真实对象
	res += this.realSubject.Do()

	// 调用之后的操作(如缓存结果，对结果进行处理等。。)
	res += ":after"

	return res
}
