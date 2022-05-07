package stack_learning_test

import (
	"math/big"
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/stack_learning"
	"github.com/holiman/uint256"
)

func TestStack(t *testing.T) {
	var stack = stack_learning.Newstack()
	var c *uint256.Int
	var data = stack.Data()
	t.Log(data)
	c = big.NewInt(1)
	stack.Push(c)
}
