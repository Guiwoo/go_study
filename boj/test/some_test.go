package test

import (
	"fmt"
	"reflect"
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

func Test_Recur(t *testing.T) {
	recur := func(n int, rs []int, start, depth, flag int) {}
	recur = func(n int, rs []int, start, depth, flag int) {
		if depth == 2 {
			fmt.Println(rs)
			return
		} else {
			for i := start; i < n; i++ {
				if (flag & (1 << i)) == 0 {
					recur(n, append(rs, i+1), i+1, depth+1, flag|(1<<i))
				}
			}
		}
	}

	recur2 := func(n int, rs []int, flag, depth int) {}
	recur2 = func(n int, rs []int, flag, depth int) {
		if depth >= 2 && len(rs) == 2 {
			fmt.Println(rs)
			return
		} else {
			for i := 0; i < n; i++ {
				if flag&(1<<i) == 0 {
					recur2(n, append(rs, i+1), flag|(1<<i), depth+1)
				}
			}
		}
	}

	recur(4, []int{}, 0, 0, 0)
	fmt.Println("======================")
	recur2(4, []int{}, 0, 0)
}

func Test_Recur_loop(t *testing.T) {
	helper := func(a, depth, start, flag, result int, out string) {}
	helper = func(a, depth, start, flag, result int, out string) {
		if depth == result {
			fmt.Println(out)
			return
		} else {
			for i := start; i < a; i++ {
				if flag&(1<<i) == 0 {
					helper(a, depth+1, i+1, flag|(1<<i), result, out+fmt.Sprintf("%d ", i))
				}
			}
		}
	}

	for i := 1; i <= 2; i++ {
		helper(4, 0, 0, 0, i, "")
		fmt.Println("===============")
	}
}

func TestBoj(t *testing.T) {
	aFunc := func(arr []int) []int {
		fmt.Printf("%p\n", arr)
		arr[0] = 9999
		return arr
	}

	arr := []int{123}

	arr2 := aFunc(arr)
	fmt.Println(arr, arr2)
	fmt.Printf("%p,%p", arr, arr2)
}

func TestEqual(t *testing.T) {
	type tester struct {
		a int
		b string
		c bool
	}
	a := tester{5, "a", true}
	b := tester{5, "a", true}
	fmt.Println("ab")
	if a == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println("cd")

	c := []int{1, 2, 3}
	d := []int{1, 2, 3}
	if reflect.DeepEqual(c, d) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
