package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)

	var size int
	fmt.Fscanln(reader, &size)

	graph := make([][]int, size)

	arr := make([][]int, size)
	for i := range arr {
		sub := make([]int, size)
		for j := range sub {
			var a int
			fmt.Fscan(reader, &a)
			if a == 1 {
				graph[i] = append(graph[i], j)
			}
			sub[j] = a
		}
		arr[i] = sub
	}

	answer := make([][]int, size)

	for i := 0; i < len(graph); i++ {
		visit := make([]bool, len(graph))
		dfs(i, graph, visit)
		sub := make([]int, size)
		for j := 0; j < len(sub); j++ {
			if visit[j] {
				sub[j] = 1
			}
		}
		answer[i] = sub
	}

	for _, v := range answer {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Println()
	}
}

func dfs(idx int, graph [][]int, visit []bool) {
	for i := 0; i < len(graph[idx]); i++ {
		if !visit[graph[idx][i]] {
			visit[graph[idx][i]] = true
			dfs(graph[idx][i], graph, visit)
		}
	}
}
