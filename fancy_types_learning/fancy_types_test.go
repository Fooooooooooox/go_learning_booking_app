package fancy_types_test

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBig(t *testing.T) {
	var v1 int64
	var v2 int64
	v1 = 100
	v2 = 101
	a := big.NewInt(v1)
	b := big.NewInt(v2)

	// big.int是没有办法直接用> <比较大小的
	// a>b返回1，a<b返回-1，a=b返回0
	TeatIsValid1 := a.Cmp(b)
	t.Log(TeatIsValid1)

	// c := big.NewInt(1)

}

func TestBigint(t *testing.T) {
	var big, whathehell = new(big.Int).SetString("2188824200011112223", 10)
	t.Log(big)
	t.Log(whathehell)
}

// 从string创建一个bigint
// 第二个参数是true
// 如果我传入的string无法被转换为bigint这个参数的值就是false
// === RUN   TestBigint
// fancy_types_test.go:27: 2188824200011112223
// fancy_types_test.go:28: true

func TestStar(t *testing.T) {

}

// iota最朴素的用法
func TestIota(t *testing.T) {
	const (
		TestA = iota
		TestB
		TestC
	)
	t.Log(TestA)
	t.Log(TestB)
	t.Log(TestC)
}

// goeth core/vm/opcodes.go 里iota的用法

func TestPrint(t *testing.T) {
	type TestOpCode byte

	const (
		// 60的十六进制是96
		PUSH1 TestOpCode = 0x60 + iota
		PUSH2
		PUSH3
		PUSH4
	)

	t.Log(PUSH1)
	t.Log(PUSH2)
	t.Log(PUSH3)
	t.Log(PUSH4)
}

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	fmt.Println(res)
	return res
}

func TestVariadic(t *testing.T) {
	Sum(1)
	Sum(1, 2, 3)
	primes := []int{2, 3, 5, 7}
	Sum(primes...)
}
