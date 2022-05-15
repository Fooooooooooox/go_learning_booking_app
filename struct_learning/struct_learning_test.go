package struct_learning_test

import (
	"testing"
	"unsafe"
)

// 看下struct类型的大小
var a uint32
var alist [3]uint32

type S struct {
	a uint16
	b uint32
}

type S2 struct {
	A struct{}
	B struct{}
}

// 为啥struct大小是8呢？==》struct的大小是把里面所有的元素大小加起来之后再加上padding（可能padding是2
// 不过如果是空的struct是没有padding的
func TestUnsafe(t *testing.T) {
	t.Log(unsafe.Sizeof(a))
	t.Log(unsafe.Sizeof(alist))
	var s S
	t.Log(unsafe.Sizeof(s))
	var s2 S2
	t.Log(unsafe.Sizeof(s2))
	// 5 is the length of the make list
	var x = make([]struct{}, 5)
	// this is a slice
	var y = x[:2]
	t.Log(x)
	t.Log(y)
}

// 空的struct能用来做啥呢？
// 1.struct{}can act as a method receiver

var set map[string]struct{}
// Initialize the set
set = make(map[string]struct{})

// Add some values to the set:
set["red"] = struct{}{}
set["blue"] = struct{}{}

// Check if a value is in the map:
_, ok := set["red"]
fmt.Println("Is red in the map?", ok)
_, ok = set["green"]
fmt.Println("Is green in the map?", ok)