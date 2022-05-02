package big_int_learning_test

import (
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
