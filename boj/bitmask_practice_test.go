package main

import (
	"fmt"
	"testing"
)

func TestBitMask01(t *testing.T) {
	rs := oddEvenChecker(124)
	fmt.Println(rs, "짝수는 true, 홀수는 false", 124)
	rs = oddEvenChecker(123)
	fmt.Println(rs, "짝수는 true,홀수는 false", 123)
}

func TestBitMask02(t *testing.T) {
	answer := toggleBitCheck(4, 4)
	fmt.Println(answer)
}

func TestBitMast3(t *testing.T) {
	arr := []bool{false, false, true}
	checkFlag(arr, 0)
	fmt.Println(arr)
}
