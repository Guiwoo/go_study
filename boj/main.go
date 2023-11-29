package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution12851()
}

func solution12851() {
	var (
		reader        = bufio.NewReader(os.Stdin)
		mover, target int
	)

	fmt.Fscanln(reader, &mover, &target)
	bfs12851(mover, target)
}

/**
5 10 9 18 17
5 4 8 16 17

0 1 2 3
0 1 2 3
*/

func bfs12851(move, target int) {
	visit := make([]int, 100001)
	for i := range visit {
		visit[i] = -1
	}
	visit[move] = 0
	path := make([]int, 100001)
	path[move] = 1
	q := []int{move}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range []int{1, -1, cur} {
			n := v + cur
			if 0 <= n && n <= 100000 {
				if visit[n] == -1 {
					visit[n] = visit[cur] + 1
					path[n] = path[cur]
					q = append(q, n)
				} else {
					if visit[n] == visit[cur]+1 {
						path[n] += path[cur]
					}
				}
			}
		}
	}
	fmt.Println(visit[target], path[target])
}
