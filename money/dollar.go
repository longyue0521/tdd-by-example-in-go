package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return &Dollar{&money{amount: amount}}
}

func (d *Dollar) Times(multiplier int) Money {
	return &Dollar{&money{amount: d.amount * multiplier}}
}
