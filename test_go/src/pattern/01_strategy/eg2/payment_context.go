package eg2

//支付上下文
type PaymentContext struct {
	Name, CardID string
	Money        int
	payment      Payment //支付
}

func NewPaymentContext(name, cardID string, money int, payment Payment) *PaymentContext {
	return &PaymentContext{
		Name:    name,
		CardID:  cardID,
		Money:   money,
		payment: payment,
	}
}

func (this *PaymentContext) Pay() {
	this.payment.Pay(this)
}
