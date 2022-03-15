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

func NewMoney(amount int, currency string) *money {
	return &money{amount: amount, currency: currency}
}

func (m *money) Amount() int {
	return m.amount
}

func (m *money) Currency() string {
	return m.currency
}

func (m *money) Equals(money Money) bool {
	return m.Amount() == money.Amount() && m.Currency() == money.Currency()
}

func (m *money) Times(multiplier int) Money {
	return nil
}
