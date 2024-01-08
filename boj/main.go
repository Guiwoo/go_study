package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boj11675()
}

func boj11675() {
	var (
		reader          = bufio.NewReader(os.Stdin)
		writer          = bufio.NewWriter(os.Stdout)
		edges, vertexes int
	)

	defer writer.Flush()
	fmt.Fscanln(reader, &edges, &vertexes)

	graph := make([][]int, 0, vertexes)
	for i := 0; i < vertexes; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)

		graph = append(graph, []int{from, to, value})
	}

	dist, ok := bellmanFord(edges, 1, graph)
	if ok {
		fmt.Println(-1)
	} else {
		for i := 2; i < len(dist); i++ {
			if dist[i] == 1e8 {
				fmt.Println(-1)
			} else {
				fmt.Println(dist[i])
			}
		}
	}
}

func bellmanFord(edges, start int, graph [][]int) ([]int, bool) {
	dp := make([]int, edges+1)
	for i := range dp {
		dp[i] = 1e8
	}
	dp[start] = 0

	for i := 0; i < edges; i++ {
		for _, t := range graph {
			from, to, val := t[0], t[1], t[2]

			if dp[from] != 1e8 && dp[to] > dp[from]+val {
				dp[to] = dp[from] + val

				if i == edges-1 {
					return dp, true
				}
			}
		}
	}
	return dp, false
}
