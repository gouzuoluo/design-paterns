package eg2

import "fmt"

//订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

/*====================================================================================================================*/

/*
* 1.关系型数据库的OrderDetailDAO实现
*/
type RDBDetailDAO struct {
}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

/*
* 2.XML存储
*/
type XMLDetailDAO struct {
}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}
