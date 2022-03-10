package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar(5)
	five.times(2)
	assert.Equal(t, 10, five.amount)
}
