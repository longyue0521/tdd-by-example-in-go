package money

type Franc struct {
	*money
	currency string
}

func NewFranc(amount int) Money {
	return &Franc{&money{amount: amount}, "CHF"}
}

func (d *Franc) Times(multiplier int) Money {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) Currency() string {
	return d.currency
}
