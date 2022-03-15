package money

type Money interface {
	Amount() int
	Currency() string
	Equals(money Money) bool
	Times(multiplier int) Money
}

type money struct {
	amount   int
	currency string
}

func (m *money) Amount() int {
	return m.amount
}

func (m *money) Currency() string {
	return m.currency
}

func (m *money) Equals(money Money) bool {
	return m.Amount() == money.Amount()
}
