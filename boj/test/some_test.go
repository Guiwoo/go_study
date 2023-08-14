package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_stringBuilder(t *testing.T) {
	sb := strings.Builder{}

	for i := 0; i < 3; i++ {
		sb.WriteString("something")
	}
	fmt.Println(sb.String())
}

func Test_bitMastk(t *testing.T) {
	fmt.Println(1 << 0)
	fmt.Println(0 | (1 << 0))
	fmt.Println(0 & (1 << 0))
}

func Test_StringSplit(t *testing.T) {
	a := "abcdefg"
	fmt.Println(a[:1])
	fmt.Println(a[1:3])
	fmt.Println(a[3:])
	fmt.Println(a[7:])
}
