package money_test

import (
	"testing"

	. "github.com/longyue0521/tdd-by-example-in-go/money"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, 10, product.Amount())
	product = five.Times(3)
	assert.Equal(t, 15, product.Amount())
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
}

func TestEqualityByEquals(t *testing.T) {
	assert.True(t, NewDollar(3).Equals(NewDollar(3)))
	assert.False(t, NewDollar(6).Equals(NewDollar(7)))
}
