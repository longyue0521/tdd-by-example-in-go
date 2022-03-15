package money

type Dollar struct {
	*money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{&money{amount: amount}}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{&money{amount: d.amount * multiplier}}
}

func (d *Dollar) Amount() int {
	return d.amount
}

func (d *Dollar) Equals(dollar *Dollar) bool {
	return d.amount == dollar.amount
}
