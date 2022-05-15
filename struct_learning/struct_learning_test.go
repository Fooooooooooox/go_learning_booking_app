package struct_learning_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
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

// var set map[string]struct{}
// // Initialize the set
// set = make(map[string]struct{})

// // Add some values to the set:
// set["red"] = struct{}{}
// set["blue"] = struct{}{}

// // Check if a value is in the map:
// _, ok := set["red"]
// fmt.Println("Is red in the map?", ok)
// _, ok = set["green"]
// fmt.Println("Is green in the map?", ok)

func TestSna(t *testing.T) {
	type Snapshot struct {
		Signers map[common.Address]struct{}
	}
	s1 := new(Snapshot)
	s1.Signers = make(map[common.Address]struct{})
	// var v1 int64
	// v1 = 100
	s1.Signers[common.HexToAddress("0x82a978b3f5962a5b0957d9ee9eef472ee55b42f1")] = struct{}{}
	t.Log(s1)
}

// 怎么在struct里用map？
// states里有cnt和category
// cnt是整数 categorty是map数组 键为string 值是自定义的events类型

// stats ==> cnt,category
// category ==>map[string]Events
// events ==> cnt, event
// event ==> value
type Stats struct {
	cnt      int
	category map[string]Events
}

// events类型有整数cnt map数组 键为string 值为自定义的event类型
type Events struct {
	cnt   int
	event map[string]*Event
}

// event类型里有value整数
type Event struct {
	value int64
}

func TestSta(t *testing.T) {
	stats := new(Stats)
	stats.cnt = 33
	stats.category = make(map[string]Events)
	e, f := stats.category["aa"]
	if !f {
		e = Events{}
	}
	e.cnt = 66
	e.event = make(map[string]*Event)
	stats.category["aa"] = e
	stats.category["aa"].event["bb"] = &Event{}
	stats.category["aa"].event["bb"].value = 99

	fmt.Println(stats)
	fmt.Println(stats.cnt, stats.category["aa"].event["bb"].value)
}
