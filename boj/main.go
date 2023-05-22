package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	params := toArr(sc)
	arrLen, _ := strconv.Atoi(params[0])
	target, _ := strconv.Atoi(params[1])

	arr := make([]int, arrLen)
	nums := toArr(sc)
	for i, numStr := range nums {
		arr[i], _ = strconv.Atoi(numStr)
	}

	answer := subArray(arr, target)
	if target == 0 {
		answer--
	}
	fmt.Println(answer)
}

func subArray(arr []int, target int) int {
	answer := 0
	n := len(arr)

	// 비트 마스크를 0부터 (2^n - 1)까지 순회합니다.
	for i := 0; i < (1 << n); i++ {
		sum := 0
		// 비트 마스크의 각 비트를 확인하여 포함된 원소의 합을 계산합니다.
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				sum += arr[j]
			}
		}
		if sum == target {
			answer++
		}
	}

	return answer
}

func toArr(sc *bufio.Scanner) []string {
	sc.Scan()
	input := sc.Text()
	rs := strings.Split(input, " ")
	return rs
}
