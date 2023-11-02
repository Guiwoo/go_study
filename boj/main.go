package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	solution()
}

// M은 20억 까지
// A의 최대값은 10억
// 1,2,3,4,5 주어진 차이 3

// 1부터 루프 돌아

func solution() {
	reader := bufio.NewReader(os.Stdin)

	var len int
	fmt.Fscanln(reader, &len)

	arr := make([]int, len)
	for i := range arr {
		var num int
		fmt.Fscan(reader, &num)
		arr[i] = num
	}

	sort.Ints(arr)

	left, right, value := 0, len-1, math.MaxInt
	index := make([]int, 2)

	for left < right {
		v := arr[left] + arr[right]
		if abs(v) < value {
			value = abs(v)
			index = []int{arr[left], arr[right]}
		} else if v > 0 {
			right--
		} else {
			left++
		}
	}

	fmt.Println(index)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
