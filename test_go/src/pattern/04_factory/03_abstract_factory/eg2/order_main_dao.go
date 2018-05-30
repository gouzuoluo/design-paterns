package eg2

import "fmt"

//订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

/*====================================================================================================================*/

/*
* 1.关系型数据库的OrderMainDAO实现
*/
type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

/*
* 2.XML存储
*/
type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}
