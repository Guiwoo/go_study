package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solution()
}

/**

1번도시의 슈퍼히어로 미노만 이동 가능 첫번쨰 줄 에는 n,m 이 주어지고
m 개의 줄에 걸쳐 x => y 로 이동할수 있는 도시의 번호를 주어진다 .

1. Create Graph with input values

2. the start point is first and should move to 6

3. so move theory twice
*/

func solution() {
	var (
		areas, edges, input int
		reader              = bufio.NewReader(os.Stdin)
		ans                 = strings.Builder{}
	)
	fmt.Fscanln(reader, &areas, &edges)
	graph := setGraph(areas, edges, reader)

	fmt.Fscanln(reader, &input)
	for i := 0; i < input; i++ {
		var bomb int
		fmt.Fscanln(reader, &bomb)
		if findWay(1, bomb, graph) && findWay(bomb, areas, graph) {
			ans.WriteString("Defend the CTP\n")
		} else {
			ans.WriteString("Destroyed the CTP\n")
		}
	}
	fmt.Println(ans.String())
}
func findWay(from, to int, graph map[int][]int) bool {
	visit := make([]bool, len(graph)+1)
	visit[from] = true
	q := []int{}
	for i := 0; i < len(graph[from]); i++ {
		visit[graph[from][i]] = true
		q = append(q, graph[from][i])
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == to {
			return true
		}
		for i := 0; i < len(graph[cur]); i++ {
			next := graph[cur][i]
			if visit[next] == false {
				q = append(q, next)
			}
		}
	}
	return false
}

func setGraph(areas, edges int, reader *bufio.Reader) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i <= areas; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Fscanln(reader, &from, &to)
		graph[from] = append(graph[from], to)
	}
	return graph
}
