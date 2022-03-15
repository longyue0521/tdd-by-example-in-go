package money

type Dollar struct {
	*money
}

func NewDollar(amount int) Money {
	return &Dollar{NewMoney(amount, "USD")}
}

func (d *Dollar) Times(multiplier int) Money {
	return &Dollar{NewMoney(d.amount*multiplier, "USD")}
}
