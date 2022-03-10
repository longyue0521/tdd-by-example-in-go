package money_test

import (
	"testing"

	. "github.com/longyue0521/tdd-by-example-in-go/money"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	assert.Equal(t, 10, five.Amount())
	five.Times(3)
	assert.Equal(t, 15, five.Amount())
}
