package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFloatTesting(t *testing.T) {
	var a float64 = 1
	for i := 0; i < 10; i++ {
		a += 0.1

		fmt.Printf("float 64 %0.20f\n", a)
	}

	b := big.NewFloat(1)
	c := big.NewFloat(0.1)

	for i := 0; i < 10; i++ {
		f := b.Add(b, c)
		fmt.Printf("big decimal %+v\n", f)
	}
}
