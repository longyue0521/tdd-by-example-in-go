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
