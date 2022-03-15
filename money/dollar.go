package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}
