package eg2

import "fmt"

//外观模块的外观接口
type FacadeAPI interface {
	Test() string
}

/*===================================================================================================================*/

//外观模式实例
type FacadeAPIImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (this *FacadeAPIImpl) Test() string {
	aRet := this.a.TestA()
	bRet := this.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewFacadeAPI() FacadeAPI {
	return &FacadeAPIImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}
