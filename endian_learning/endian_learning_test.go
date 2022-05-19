package endian_learning_test

import (
	"testing"
)

// func TestEn(t *testing.T) {
// 	n := uint32(0x01020304)
// 	t.Log(n)
// 	t.Log(unsafe.Pointer(&n))
// 	t.Log((*byte)(unsafe.Pointer(&n)))
// 	t.Log(*(*byte)(unsafe.Pointer(&n)))
// }

func TestByte(t *testing.T) {
	n := uint32(900)
	var b byte = byte(n)
	t.Log(b)
	t.Log(n)
}
