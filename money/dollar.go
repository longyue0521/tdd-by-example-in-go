package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

func (d *Dollar) Times(multiplier int) Money {
	return NewMoney(d.amount*multiplier, d.currency)
}
