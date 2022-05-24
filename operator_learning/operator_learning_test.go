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
