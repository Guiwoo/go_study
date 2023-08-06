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

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	boj2529(arr)

}

var minResult, maxResult string

func boj2529(arr []string) {
	maxResult, minResult = "", "999999999"
	for i := 0; i <= 9; i++ {
		boj2529Recur(arr, []int{i}, 0, 1<<i)
	}
	fmt.Println(maxResult)
	fmt.Println(minResult)
}

func boj2529Recur(arr []string, result []int, idx, flag int) {
	if len(result) == len(arr)+1 {
		var str string
		for _, v := range result {
			str += fmt.Sprintf("%d", v)
		}
		if str > maxResult {
			maxResult = str
		}
		if str < minResult {
			minResult = str
		}
		return
	} else {
		if arr[idx] == "<" {
			for i := result[idx] + 1; i <= 9; i++ {
				if flag&(1<<i) == 0 {
					boj2529Recur(arr, append(result, i), idx+1, flag|(1<<i))
				}
			}
		} else {
			for i := result[idx] - 1; i >= 0; i-- {
				if flag&(1<<i) == 0 {
					boj2529Recur(arr, append(result, i), idx+1, flag|(1<<i))
				}
			}
		}
	}
}
