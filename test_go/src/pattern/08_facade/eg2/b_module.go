package eg2

type BModuleAPI interface {
	TestB() string
}

/*==================================================================================================================*/
//B模块实例
type BModuleImpl struct {
}

func (*BModuleImpl) TestB() string {
	return "B module running"
}

func NewBModuleAPI() BModuleAPI {
	return &BModuleImpl{}
}
