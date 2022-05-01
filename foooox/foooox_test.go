package foooox_test

import (
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/foooox"
)

func TestFoooox(t *testing.T) {
	if foooox.Bestfoooox() != "hhh" {
		t.Fatal("wrong foooox")
	}
}
