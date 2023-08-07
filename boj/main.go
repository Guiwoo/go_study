package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	oper := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &oper[i])
	}
	boj14888(arr, oper)
}

var max = -1000000000
var min = 1000000000

// + - * % 개수
func boj14888(arr, oper []int) {
	boj14888Recur(arr, oper, 0, arr[0])
	fmt.Println(max, min)
}
func boj14888Recur(arr, oper []int, idx, rs int) {
	if idx == len(arr)-1 {
		if rs > max {
			max = rs
		}
		if rs < min {
			min = rs
		}
		return
	} else {
		for i := 0; i < len(oper); i++ {
			switch i {
			case 0:
				if oper[i] > 0 {
					oper[i]--
					boj14888Recur(arr, oper, idx+1, rs+arr[idx+1])
					oper[i]++
				}
			case 1:
				if oper[i] > 0 {
					oper[i]--
					boj14888Recur(arr, oper, idx+1, rs-arr[idx+1])
					oper[i]++
				}
			case 2:
				if oper[i] > 0 {
					oper[i]--
					boj14888Recur(arr, oper, idx+1, rs*arr[idx+1])
					oper[i]++
				}
			default:
				if oper[i] > 0 {
					oper[i]--
					boj14888Recur(arr, oper, idx+1, rs/arr[idx+1])
					oper[i]++
				}
			}
		}
	}
}
