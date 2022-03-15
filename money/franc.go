package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return &Franc{&money{amount: amount, currency: "CHF"}}
}

func (d *Franc) Times(multiplier int) Money {
	return NewFranc(d.amount * multiplier)
}
