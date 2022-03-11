package money

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

func (d *Franc) Times(multiplier int) *Franc {
	return &Franc{amount: d.amount * multiplier}
}

func (d *Franc) Equals(dollar *Franc) bool {
	return d.amount == dollar.amount
}
