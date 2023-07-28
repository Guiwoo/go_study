package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var buf bytes.Buffer

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
	boj15657(arr, m)
}

func boj15657(n []int, m int) {
	out := make([]int, m)
	boj15657Recur(n, out, 0)
	fmt.Println(buf.String())
}
func boj15657Recur(n, out []int, depth int) {
	if depth == len(out) {
		for _, v := range out {
			buf.WriteString(strconv.Itoa(v) + " ")
		}
		buf.WriteByte('\n')
		return
	}

	for i := 0; i < len(n); i++ {
		if depth == 0 || out[depth-1] <= n[i] {
			out[depth] = n[i]
			boj15657Recur(n, out, depth+1)
		}
	}
}
