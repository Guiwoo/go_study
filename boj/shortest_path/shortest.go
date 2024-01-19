package shortest_path

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
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

func boj9370() {
	var (
		reader     = bufio.NewReader(os.Stdin)
		writer     = bufio.NewWriter(os.Stdout)
		simulation int
	)
	fmt.Fscanln(reader, &simulation)
	defer writer.Flush()

	for i := 0; i < simulation; i++ {
		answer := lookingForHeadingTo(reader)
		writer.WriteString(answer + "\n")
	}
}

func lookingForHeadingTo(reader *bufio.Reader) string {
	var (
		vertexes, edges, candidates int
		start, cur1, cur2           int
	)

	fmt.Fscanln(reader, &vertexes, &edges, &candidates)
	fmt.Fscanln(reader, &start, &cur1, &cur2)

	graph := make([]map[int]int, vertexes+1)

	for i := range graph {
		graph[i] = make(map[int]int)
	}

	for i := 0; i < edges; i++ {
		var (
			from, to, value int
		)
		fmt.Fscanln(reader, &from, &to, &value)
		graph[from][to] = value
		graph[to][from] = value
	}

	candArr := make([]int, candidates)
	for i := 0; i < candidates; i++ {
		var x int
		fmt.Fscanln(reader, &x)
		candArr[i] = x
	}

	all := dijkstra9370(graph, start, vertexes)
	startToFirst := dijkstra9370(graph, cur1, vertexes)
	startToSecond := dijkstra9370(graph, cur2, vertexes)

	answer := make([]int, 0)
	for _, v := range candArr {

		if all[v] == all[cur1]+startToFirst[cur2]+startToSecond[v] ||
			all[v] == all[cur2]+startToSecond[cur1]+startToFirst[v] {
			answer = append(answer, v)
			continue
		}
	}
	sort.Ints(answer)
	str := ""
	for _, v := range answer {
		str += fmt.Sprintf("%d ", v)
	}
	return str
}

func dijkstra9370(graph []map[int]int, start int, vertexes int) []int {
	dp := make([]int, vertexes+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	var pq PriorityQ
	heap.Init(&pq)

	dp[start] = 0
	heap.Push(&pq, &item{cur: start, value: 0})

	for pq.Len() > 0 {
		current := pq.Pop().(*item)

		for next, value := range graph[current.cur] {
			nextValue := current.value + value
			if nextValue < dp[next] {
				dp[next] = nextValue
				heap.Push(&pq, &item{next, nextValue})
			}
		}
	}

	return dp
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
		if !relaxd(dp, graph, edges) {
			fmt.Println("NO")
			return
		}
	}
	if !relaxd(dp, graph, edges) {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	return
}

func relaxd(dp []int, graph []map[int][]int, edges int) bool {
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
