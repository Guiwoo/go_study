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

	result := boj15650(n, m)
	writer.WriteString(result)
}

func boj15650(n, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	recur(n, 0, 0, 0, out, &sb)
	return sb.String()
}

func recur(n, start, depth, flag int, out []int, sb *strings.Builder) {
	if depth == len(out) {
		result := strings.NewReplacer("[", "", "]", "").Replace(fmt.Sprintf("%v", out))
		sb.WriteString(result + "\n")
		return
	} else {
		for i := start; i < n; i++ {
			if flag&(1<<i) != 0 {
				continue
			}
			out[depth] = i + 1
			recur(n, i+1, depth+1, flag|(1<<i), out, sb)
		}
	}
}
