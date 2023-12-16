package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	boj1753()
}

/*
*
방향그래프가 주어지면 주어진 시작점에서 다른 모든 정점으로의 최단 경로를 구하는 프로그램을 작성하시오. 단, 모든 간선의 가중치는 10 이하의 자연수이다.

첫째 줄에 정점의 개수 V와 간선의 개수 E가 주어진다.(1 ≤ V ≤ 20,000, 1 ≤ E ≤ 300,000)
모든 정점에는 1부터 V까지 번호가 매겨져 있다고 가정한다.
둘째 줄에는 시작 정점의 번호 K(1 ≤ K ≤ V)가 주어진다.
셋째 줄부터 E개의 줄에 걸쳐 각 간선을 나타내는 세 개의 정수 (u, v, w)가 순서대로 주어진다.
이는 u에서 v로 가는 가중치 w인 간선이 존재한다는 뜻이다. u와 v는 서로 다르며 w는 10 이하의 자연수이다.
서로 다른 두 정점 사이에 여러 개의 간선이 존재할 수도 있음에 유의한다.

첫째 줄부터 V개의 줄에 걸쳐,
i번째 줄에 i번 정점으로의 최단 경로의 경로값을 출력한다.
시작점 자신은 0으로 출력하고, 경로가 존재하지 않는 경우에는 INF를 출력하면 된다.

5 6
1
5 1 1
1 2 2
1 3 3
2 3 4
2 4 5
3 4 6

0
2
3
7
INF
*/
type item struct {
	cur, value int
	index      int
}
type PriorityQ []*item

func (pq PriorityQ) Len() int {
	return len(pq)
}
func (pq PriorityQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}
func (pq PriorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQ) Push(x any) {
	item := x.(*item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *PriorityQ) Pop() any {
	old := *pq
	item := old[0]
	old[0] = nil
	item.index = -1
	*pq = old[1:]
	return item
}

func boj1753() {

	var (
		reader              = bufio.NewReader(os.Stdin)
		writer              = bufio.NewWriter(os.Stdout)
		vertex, edge, start int
	)

	defer writer.Flush()
	fmt.Fscanln(reader, &vertex, &edge)
	fmt.Fscanln(reader, &start)

	graph := make([]map[int]int, vertex+1)
	for i := range graph {
		graph[i] = make(map[int]int, 0)
	}

	for i := 0; i < edge; i++ {
		var (
			inputFrom, inputTo, value int
		)
		fmt.Fscanln(reader, &inputFrom, &inputTo, &value)
		if v, ok := graph[inputFrom][inputTo]; !ok || v > value {
			graph[inputFrom][inputTo] = value
		}
	}

	check := make([]int, vertex+1)
	for i := range check {
		check[i] = math.MaxInt
	}
	q := make(PriorityQ, 0)
	q.Push(&item{start, 0, 0})
	heap.Init(&q)
	check[start] = 0
	for q.Len() > 0 {
		current := q.Pop()
		cur := current.(*item)
		if check[cur.cur] < cur.value {
			continue
		}
		for to, value := range graph[cur.cur] {
			if check[to] > check[cur.cur]+value {
				check[to] = check[cur.cur] + value
				heap.Push(&q, &item{to, check[to], 0})
			}
		}
	}
	sb := strings.Builder{}
	for i := 1; i <= vertex; i++ {
		if check[i] == math.MaxInt {
			sb.WriteString("INF\n")
		} else {
			sb.WriteString(fmt.Sprintf("%d\n", check[i]))
		}
	}
	fmt.Fprintln(writer, sb.String())
}
