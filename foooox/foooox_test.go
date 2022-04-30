package foooox_test

import (
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app"
)

func TestFoooox(t *testing.T) {
	if go_learning_booking_app.Bestfoooox() != "hhh" {
		t.Fatal("wrong foooox")
	}
}
