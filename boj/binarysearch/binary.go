package binarysearch

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func boj3273() {
	var (
		reader       = bufio.NewReader(os.Stdin)
		writer       = bufio.NewWriter(os.Stdout)
		N, T, answer int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	arr := make([]int, N)

	input, _ := reader.ReadString('\n')
	for i, v := range strings.Split(strings.TrimSpace(input), " ") {
		x, _ := strconv.Atoi(v)
		arr[i] = x
	}

	fmt.Fscanln(reader, &T)

	sort.Ints(arr)

	start, end := 0, len(arr)-1

	for start < end {
		val := arr[start] + arr[end]
		if val > T {
			end--
		} else if val < T {
			start++
		} else if val == T {
			start, end, answer = start+1, end-1, answer+1
		}
	}

	fmt.Fprintln(writer, answer)
}
