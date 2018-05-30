package eg2

//静态工厂方法
func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	} else if t == 2 {
		return &HelloAPI{}
	}
	return nil
}
