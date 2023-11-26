package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	areas, edges, input int
	reader              = bufio.NewReader(os.Stdin)
	ans                 = strings.Builder{}
	visit               = [2][100001]bool{}
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

	fmt.Fscanln(reader, &areas, &edges)
	graph, reverse := setGraph(areas, edges, reader)

	findWay(1, areas, 0, graph)
	findWay(areas, 1, 1, reverse)

	fmt.Fscanln(reader, &input)
	for i := 0; i < input; i++ {
		var bomb int
		fmt.Fscanln(reader, &bomb)
		if visit[0][bomb] && visit[1][bomb] {
			ans.WriteString("Defend the CTP\n")
		} else {
			ans.WriteString("Destroyed the CTP\n")
		}
	}
	fmt.Println(ans.String())
}
func findWay(from, to, way int, graph map[int][]int) {
	visit[way][from] = true
	q := []int{from}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < len(graph[cur]); i++ {
			next := graph[cur][i]
			if visit[way][next] == false {
				visit[way][next] = true
				q = append(q, next)
			}
		}
	}
	return
}

func setGraph(areas, edges int, reader *bufio.Reader) (map[int][]int, map[int][]int) {
	graph := make(map[int][]int)
	reverse := make(map[int][]int)
	for i := 0; i <= areas; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Fscanln(reader, &from, &to)
		graph[from] = append(graph[from], to)
		reverse[to] = append(reverse[to], from)
	}
	return graph, reverse
}
