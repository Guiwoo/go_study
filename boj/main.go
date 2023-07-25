package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(input, " ")
	n, _ := strconv.Atoi(inputs[0])
	m, _ := strconv.Atoi(strings.TrimSpace(inputs[1]))
	//writer.WriteString(fmt.Sprintf("%d  %d", n, m))
	writer.WriteString(boj15652(n, m))
}

func boj15652(n, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	permutation(n, 0, out, &sb)
	return sb.String()
}
func permutation(n, depth int, out []int, sb *strings.Builder) {
	if depth == len(out) {
		for _, v := range out {
			sb.WriteString(fmt.Sprintf("%d ", v))
		}
		sb.WriteString("\n")
		return
	} else {
		start := 0
		if depth != 0 {
			start = out[depth-1] - 1
		}
		for i := start; i < n; i++ {
			out[depth] = i + 1
			permutation(n, depth+1, out, sb)
		}
	}
}
