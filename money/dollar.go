package money

type Dollar struct {
	*money
	currency string
}

func NewDollar(amount int) Money {
	return &Dollar{&money{amount: amount}, "USD"}
}

func (d *Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) Currency() string {
	return d.currency
}
