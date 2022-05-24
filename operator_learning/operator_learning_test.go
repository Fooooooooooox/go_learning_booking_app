package operatorlearning_test

import (
	"fmt"
	"testing"
)

func TestOperators(t *testing.T) {
	var x, y uint
	x, y = 1, 1

	for x = 1; x < 10; x++ {
		fmt.Printf("%d << %d = %d \n", x, y, x<<y)
	}
	fmt.Println()

	x = 512
	for y = 1; y < 10; y++ {
		fmt.Printf("%d >> %d = %d \n", x, y, x>>y)
	}

	// n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

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
