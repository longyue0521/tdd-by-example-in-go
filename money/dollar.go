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
