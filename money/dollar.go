package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return newDollar(amount, "USD")
}

func (d *Dollar) Times(multiplier int) Money {
	return newDollar(d.amount*multiplier, d.currency)
}

func newDollar(amount int, currency string) Money {
	return &Dollar{NewMoney(amount, currency)}
}
