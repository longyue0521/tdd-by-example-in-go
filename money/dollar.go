package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return &Dollar{&money{amount: amount, currency: "USD"}}
}

func (d *Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}
