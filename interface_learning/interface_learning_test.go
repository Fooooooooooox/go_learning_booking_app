package interface_learning_test

import (
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/interface_learning"
)

func TestInterface(t *testing.T) {
	TestRes := interface_learning.Interface()
	t.Log(TestRes)
}
