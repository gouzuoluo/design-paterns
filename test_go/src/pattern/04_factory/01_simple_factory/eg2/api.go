package eg2

import "fmt"

//产品接口
type API interface {
	Say(name string) string
}

/*===================================================================================================================*/

/*
* 1.实现1
*/
type HiAPI struct {
}

func (*HiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

/*
* 2.实现2
*/
type HelloAPI struct {
}

func (*HelloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
