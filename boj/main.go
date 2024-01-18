package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boj1865()
}

func boj1865() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		tcs    int
	)

	defer writer.Flush()
	fmt.Fscanln(reader, &tcs)

	for i := 0; i < tcs; i++ {
		isBlackHole(reader)
	}
}

func isBlackHole(reader *bufio.Reader) {
	var (
		edges, vertexes, wormhole int
	)

	fmt.Fscanln(reader, &edges, &vertexes, &wormhole)

	graph := make([]map[int][]int, edges+1)
	for i := range graph {
		graph[i] = make(map[int][]int)
	}

	for i := 0; i < vertexes; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)
		graph[from][to] = append(graph[from][to], value)
		graph[to][from] = append(graph[to][from], value)
	}

	for i := 0; i < wormhole; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)
		graph[from][to] = append(graph[from][to], -value)
	}

	dp := make([]int, edges+1)
	for i := range dp {
		dp[i] = 1e8
	}

	for i := 1; i < edges; i++ {
		if !relax(dp, graph, edges) {
			fmt.Println("NO")
			return
		}
	}
	if !relax(dp, graph, edges) {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	return
}

func relax(dp []int, graph []map[int][]int, edges int) bool {
	relaxed := false
	for from := 1; from <= edges; from++ {
		for to, values := range graph[from] {
			for _, value := range values {
				if dp[to] > dp[from]+value {
					dp[to] = dp[from] + value
					relaxed = true
				}
			}
		}
	}
	return relaxed
}
