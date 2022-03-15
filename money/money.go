package money

type Money interface {
	Amount() int
	Equals(money Money) bool
}

type money struct {
	amount int
}

func (m *money) Amount() int {
	return m.amount
}

func (m *money) Equals(money Money) bool {
	return m.Amount() ==  money.Amount()
}