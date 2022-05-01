package json_learning_test

import (
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/json_learning"
)

func TestStudent(t *testing.T) {
	Res := json_learning.Students()
	t.Log(Res)
}
