package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return &Franc{NewMoney(amount, "CHF")}
}

func (d *Franc) Times(multiplier int) Money {
	return &Franc{NewMoney(d.amount*multiplier, d.currency)}
}

func newFranc(amount int, currency string) Money {
	return &Franc{NewMoney(amount, currency)}
}
