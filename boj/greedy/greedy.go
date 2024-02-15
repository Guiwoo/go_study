package greedy

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func boj11047() {
	var (
		reader            = bufio.NewReader(os.Stdin)
		writer            = bufio.NewWriter(os.Stdout)
		N, Target, answer int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &N, &Target)

	s := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscanln(reader, &x)
		s = append(s, x)
	}

	for i := N - 1; i >= 0; i-- {
		if s[i] > Target || Target == 0 {
			continue
		}
		if s[i] == Target {
			Target -= s[i]
			answer++
			break
		}
		if s[i] < Target {
			x := Target / s[i]
			Target -= s[i] * x
			answer += x
		}
	}

	fmt.Fprintln(writer, answer)
}

type Meeting struct {
	start, end int
}

func boj1931() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		N      int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	arr := make([]Meeting, 0, N)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		arr = append(arr, Meeting{x, y})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].end == arr[j].end {
			return arr[i].start < arr[j].start
		}
		return arr[i].end < arr[j].end
	})

	var (
		tmp, answer = 0, 0
	)

	for i := 0; i < len(arr); i++ {
		if tmp <= arr[i].start {
			tmp = arr[i].end
			answer++
		}
	}

	fmt.Fprintln(writer, answer)
}
