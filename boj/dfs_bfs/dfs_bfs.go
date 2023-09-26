package dfs_bfs

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func boj1260() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	arr := make([][]int, a+1)
	edge := 0
	for i := 0; i < b; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		arr[x] = append(arr[x], y)
		arr[y] = append(arr[y], x)
		edge++
	}

	for i := range arr {
		sort.Ints(arr[i])
	}
	var (
		dfs_answer []int
		bfs_answer []int
	)

	dfs := func(cur int, visit []bool) {}
	dfs = func(cur int, visit []bool) {
		target := arr[cur]
		for i := 0; i < len(target); i++ {
			if !visit[target[i]] {
				visit[target[i]] = true
				dfs_answer = append(dfs_answer, target[i])
				dfs(target[i], visit)
			}
		}
	}

	bfs := func(start int) {
		q := []int{start}
		visit := make([]bool, a+1)
		for len(q) > 0 {
			target := q[0]
			q = q[1:]
			if visit[target] {
				continue
			}
			visit[target] = true
			bfs_answer = append(bfs_answer, target)
			for i := 0; i < len(arr[target]); i++ {
				next := arr[target][i]
				if !visit[next] {
					q = append(q, next)
				}
			}
		}
	}
	visit := make([]bool, a+1)
	visit[c] = true

	dfs_answer = append(dfs_answer, c)
	dfs(c, visit)
	bfs(c)

	answer := strings.Builder{}
	for _, v := range dfs_answer {
		answer.WriteString(strconv.Itoa(v) + " ")
	}
	answer.WriteString("\n")
	for _, v := range bfs_answer {
		answer.WriteString(strconv.Itoa(v) + " ")
	}

	fmt.Println(answer.String())
}

func boj11724() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	arr := make([][]int, a+1)

	for i := 0; i < b; i++ {
		var start, end int
		fmt.Fscanln(reader, &start, &end)
		arr[start] = append(arr[start], end)
		arr[end] = append(arr[end], start)
	}
	visit := make([]bool, a+1)
	ans := 0
	for i := 1; i <= a; i++ {
		if !visit[i] {
			visit[i] = true
			bfs_11724(i, arr, visit)
			ans++
		}
	}
	fmt.Println(ans)
}

// dfs 가 조금더 빠른것으로 보임 왜그런지 확인하고 코드 작성해보기
func bfs_11724(start int, arr [][]int, visit []bool) {
	q := []int{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for i := 0; i < len(arr[cur]); i++ {
			if !visit[arr[cur][i]] {
				visit[arr[cur][i]] = true
				check := false
				for _, v := range arr[arr[cur][i]] {
					if !visit[v] {
						check = true
						break
					}
				}
				if check {
					q = append(q, arr[cur][i])
				}
			}

		}
	}
}
