package shortest_path

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type item struct {
	cur, value int
}

type PriorityQ []*item

func (pq *PriorityQ) Len() int {
	q := *pq
	return len(q)
}
func (pq *PriorityQ) Swap(i, j int) {
	q := *pq
	q[i], q[j] = q[j], q[i]
}
func (pq *PriorityQ) Less(i, j int) bool {
	q := *pq
	return q[i].value < q[j].value
}

func (pq *PriorityQ) Push(x any) {
	item := x.(*item)
	*pq = append(*pq, item)
}
func (pq *PriorityQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func boj1504() {
	var (
		reader                = bufio.NewReader(os.Stdin)
		writer                = bufio.NewWriter(os.Stdout)
		n, edges              int
		edgeFirst, edgeSecond int
	)
	defer writer.Flush()
	fmt.Fscanln(reader, &n, &edges)

	graph := make([]map[int]int, n+1)

	for i := range graph {
		graph[i] = make(map[int]int)
	}

	for i := 0; i < edges; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)
		m := graph[from]
		m[to] = value
		rev := graph[to]
		rev[from] = value
	}

	fmt.Fscanln(reader, &edgeFirst, &edgeSecond)

	startToFirst := dijkstra1504(1, edgeFirst, graph)
	firstToSecond := dijkstra1504(edgeFirst, edgeSecond, graph)
	secondToEnd := dijkstra1504(edgeSecond, n, graph)

	startToSecond := dijkstra1504(1, edgeSecond, graph)
	secondToFirst := dijkstra1504(edgeSecond, edgeFirst, graph)
	firstToEnd := dijkstra1504(edgeFirst, n, graph)

	firstWay := startToSecond + secondToFirst + firstToEnd
	secondWay := startToFirst + firstToSecond + secondToEnd
	if startToSecond == math.MaxInt || secondToFirst == math.MaxInt || firstToEnd == math.MaxInt {
		firstWay = math.MaxInt
	}
	if startToFirst == math.MaxInt || firstToSecond == math.MaxInt || secondToEnd == math.MaxInt {
		secondWay = math.MaxInt
	}
	ans := min(firstWay, secondWay)
	if ans < 0 || ans >= math.MaxInt {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func dijkstra1504(from, to int, graph []map[int]int) int {
	visit := make([]int, len(graph)+1)
	for i := range visit {
		visit[i] = math.MaxInt
	}

	pq := make(PriorityQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, &item{from, 0})
	visit[from] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*item)
		for next, add := range graph[cur.cur] {
			if visit[next] > cur.value+add {
				visit[next] = cur.value + add
				heap.Push(&pq, &item{next, visit[next]})
			}
		}
	}
	return visit[to]
}
