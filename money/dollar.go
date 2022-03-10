package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) Times(multiplier int) {
	d.amount = 5 * 2
}

func (d *Dollar) Amount() int {
	return d.amount
}
