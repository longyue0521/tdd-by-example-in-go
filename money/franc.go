package money

type Franc struct {
	*money
}

func NewFranc(amount int) *Franc {
	return &Franc{&money{amount: amount}}
}

func (d *Franc) Times(multiplier int) *Franc {
	return &Franc{&money{amount: d.amount * multiplier}}
}

func (d *Franc) Amount() int {
	return d.amount
}

func (d *Franc) Equals(dollar *Franc) bool {
	return d.amount == dollar.amount
}
