package stringmemorylearning_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
)

func TestJianyu(t *testing.T) {
	s := "脑子进煎鱼了"
	s1 := "脑子进煎鱼了"
	s2 := "脑子进煎鱼了"[7:]

	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data)
	t.Log(common.BytesToAddress([]byte{1}))
}

// go 中的内存管理 这三个string指向的内存都是一样的
