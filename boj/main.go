package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	solution13913()
}

var (
	reader   = bufio.NewReader(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	from, to int

	visit [100001]int
	path  [100001]int
)

func solution13913() {

	defer writer.Flush()
	fmt.Fscan(reader, &from, &to)

	q := list.New()

	visit[from] = 1
	path[from] = from
	q.PushBack(from)

	for q.Len() > 0 {
		cur := q.Front().Value.(int)
		q.Remove(q.Front())

		if cur == to {
			break
		}
		for _, v := range []int{-1, cur, 1} {
			next := v + cur
			if next >= 0 && next <= 100000 && visit[next] == 0 {
				visit[next] = visit[cur] + 1
				path[next] = cur
				q.PushBack(next)
			}
		}
	}

	fmt.Fprintln(writer, visit[to]-1)
	getPath(to)
}

func getPath(u int) {
	if u == path[u] {
		fmt.Fprint(writer, u, " ")
		return
	}
	getPath(path[u])
	fmt.Fprint(writer, u, " ")
}
