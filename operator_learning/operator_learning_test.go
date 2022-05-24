package operatorlearning_test

import (
	"fmt"
	"testing"
)

func TestOperators(t *testing.T) {
	var x, y uint
	x, y = 1, 2

	for x = 1; x < 10; x++ {
		fmt.Printf("%d << %d = %d \n", x, y, x<<y)
	}
	fmt.Println()

	x = 512
	for y = 1; y < 10; y++ {
		fmt.Printf("%d >> %d = %d \n", x, y, x>>y)
	}

	// n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

	// 2的9次方乘1 = 512
	// 2的8次方乘2 = 512
	// 2的7次方乘4 = 512

	// 所以<< >> 这个东西就是和2的指数相关的操作
	// === RUN   TestOperators
	// 1 << 2 = 4
	// 2 << 2 = 8
	// 3 << 2 = 12
	// 4 << 2 = 16
	// 5 << 2 = 20
	// 6 << 2 = 24
	// 7 << 2 = 28
	// 8 << 2 = 32
	// 9 << 2 = 36

	// 512 >> 1 = 256
	// 512 >> 2 = 128
	// 512 >> 3 = 64
	// 512 >> 4 = 32
	// 512 >> 5 = 16
	// 512 >> 6 = 8
	// 512 >> 7 = 4
	// 512 >> 8 = 2
	// 512 >> 9 = 1

}

// bitwiseor
// |= 和 +=是一个意思
// Like a += b is equivalent to a = a + b, a |= b is equivalent to a = a | b.

func TestBitwiseor(t *testing.T) {
	var x, y uint
	x, y = 1, 0
	y |= x
	t.Log(x)
	t.Log(y)
}
