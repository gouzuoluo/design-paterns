package eg2

import "fmt"

type Payment interface {
	Pay(*PaymentContext)
}

/*====================================================================================================================*/
//1.现金支付
type Cash struct {
}

func (*Cash) Pay(pc *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", pc.Money, pc.Name)
}

//2.银行卡支付
type Bank struct {
}

func (*Bank) Pay(pc *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", pc.Money, pc.Name, pc.CardID)

}
