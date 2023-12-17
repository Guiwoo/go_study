package main

import (
	"container/heap"
	"fmt"
)

func main() {
	boj1504()
}

/**
1. 방향성이 없는 그래프 에서 1 ~ N 번으로 이동하려고 한다.
2. 임의로 주어진 두 정점을 반드시 통과해야 한다.
3. 이동했던 간선도 이동 가능하다.
4. 반드시 최단 경ㄹ로로 이동 해야 한다.

*/

type Item struct {
	heading, value int
}
type PriorityQ []*Item

func (pq PriorityQ) Len() int {
	return len(pq)
}
func (pq PriorityQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}
func (pq PriorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQ) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func boj1504() {
	q := make(PriorityQ, 0)
	heap.Init(&q)

	heap.Push(&q, &Item{heading: 1, value: 9})
	heap.Push(&q, &Item{heading: 4, value: 2})
	heap.Push(&q, &Item{heading: 5, value: 1})
	heap.Push(&q, &Item{heading: 6, value: 3})
	heap.Push(&q, &Item{heading: 7, value: 0})

	for q.Len() > 0 {
		a := heap.Pop(&q)
		fmt.Printf("%+v\n", a)
	}
}
