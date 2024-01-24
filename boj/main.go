package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boj1956()
}

func boj1956() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		N, V   int
		answer int = 1e8
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &V)

	graph := make([][]int, N+1)
	for i := range graph {
		sub := make([]int, N+1)
		for j := range sub {
			//if i == j {
			//	continue
			//}
			sub[j] = 1e8
		}
		graph[i] = sub
	}

	for i := 0; i < V; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)

		graph[from][to] = value
	}

	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}

	for i := 1; i <= N; i++ {
		answer = min(answer, graph[i][i])
	}

	if answer == 1e8 {
		answer = -1
	}

	fmt.Fprintln(writer, answer)
}
