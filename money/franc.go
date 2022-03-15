package money

type Franc struct {
	*money
}

func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
