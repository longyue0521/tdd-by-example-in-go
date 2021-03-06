package money_test

import (
	"testing"

	. "github.com/longyue0521/tdd-by-example-in-go/money"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.Equal(t, NewFranc(2), NewFranc(2))
}

func TestEqualityByEquals(t *testing.T) {
	assert.True(t, NewDollar(3).Equals(NewDollar(3)))
	assert.False(t, NewDollar(6).Equals(NewDollar(7)))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.Times(2))
	assert.Equal(t, NewFranc(15), five.Times(3))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestDifferentTypeEquality(t *testing.T) {
	assert.False(t, NewDollar(10).Equals(NewFranc(10)))
}

func TestSimpleAddition(t *testing.T) {
	sum := NewDollar(5).Plus(NewDollar(5))
	assert.True(t, NewDollar(10).Equals(sum))
	assert.Equal(t, NewDollar(10), sum)
}
