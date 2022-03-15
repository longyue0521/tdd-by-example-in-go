package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return &Dollar{&money{amount: amount}}
}

func (d *Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) Currency() string {
	return "USD"
}
