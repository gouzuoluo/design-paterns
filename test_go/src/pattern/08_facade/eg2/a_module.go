package eg2

//A模块API
type AModuleAPI interface {
	TestA() string
}

/*===================================================================================================================*/

//A模块实例
type AModuleImpl struct {
}

func (*AModuleImpl) TestA() string {
	return "A module running"
}

func NewAModuleAPI() AModuleAPI {
	return &AModuleImpl{}
}
