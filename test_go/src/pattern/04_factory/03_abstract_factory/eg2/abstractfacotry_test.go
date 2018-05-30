package eg2

import "testing"

func TestAll(t *testing.T) {
	var (
		factory DAOFactory
	)

	//1.RDB工厂测试
	factory = &RDBDAOFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

	//2.XML工厂测试
	factory = &XMLDAOFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}
