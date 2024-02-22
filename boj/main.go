package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
func main(){
	fmt.Fscanln(reader, &N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(reader, &x)
		arr[i] = x
	}

	sort.Ints(arr)

	dp, sum := make([]int, N), arr[0]
	dp[0] = arr[0]

	for i := 1; i < N; i++ {
		dp[i] = arr[i] + dp[i-1]
		sum += dp[i]

	}

	fmt.Fprintln(writer, sum)
}
