package float

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_Float(t *testing.T) {
	var x float64 = 0
	for i := 0; i < 10; i++ {
		x += 0.1
	}

	fmt.Println(x == 1)
}

func Test_BigDecimal(t *testing.T) {
	n := big.NewFloat(0)

	for i := 0; i < 10; i++ {
		n.Add(n, big.NewFloat(0.1))
	}

	fmt.Println(n)

}
