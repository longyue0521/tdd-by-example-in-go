package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return &Franc{&money{amount: amount}}
}

func (d *Franc) Times(multiplier int) Money {
	return &Franc{&money{amount: d.amount * multiplier}}
}
