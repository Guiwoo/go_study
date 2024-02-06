package greedy

import (
	"bufio"
	"fmt"
	"os"
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
