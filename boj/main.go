package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	solution()
}

/**
5 5
1 3
1 4
4 5
4 3
3 2

3


임의의 두사람이 최소 몇단계 만에 이어질수 있는가 계산하는 법칙
모든 사람에게 갈수 있는 경로를 파악하고 ++
*/

func solution() {
	var (
		reader           = bufio.NewReader(os.Stdin)
		users, relations int
		index            int
	)

	fmt.Fscanln(reader, &users, &relations)
	graph := make([][]int, users+1)
	answer := make([]int, users+1)
	for i := range answer {
		answer[i] = math.MaxInt
	}

	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < relations; i++ {
		var who, know int
		fmt.Fscanln(reader, &who, &know)
		graph[who] = append(graph[who], know)
		graph[know] = append(graph[know], who)
	}

	for user := 1; user <= users; user++ {
		var (
			cnt   = 0
			q     = []int{user}
			visit = make([]bool, users+1)
			run   = 0
		)

		visit[user] = true

		for len(q) > 0 {
			size := len(q)
			for i := 0; i < size; i++ {
				cur := q[0]
				cnt += run
				q = q[1:]
				for _, v := range graph[cur] {
					if visit[v] == false {
						visit[v] = true
						q = append(q, v)
					}
				}
			}
			run++
		}

		if answer[index] > cnt {
			index = user
			answer[index] = cnt
		}
	}

	fmt.Println(index)
}
