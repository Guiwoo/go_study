package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)
	writer.WriteString()
}
