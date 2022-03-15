package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}

func (d *Franc) Times(multiplier int) Money {
	return NewMoney(d.amount*multiplier, d.currency)
}

