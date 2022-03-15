package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return &Franc{&money{amount: amount}}
}

func (d *Franc) Times(multiplier int) Money {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) Currency() string {
	return "CHF"
}