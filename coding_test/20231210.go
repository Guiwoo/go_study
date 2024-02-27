package main

import "container/list"

func solution5(board [][]int, c int) int {
	var (
		dirs          = []int{0, 1, 0, -1, 0}
		start, target []int
	)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 2 {
				start = []int{i, j}
			}
			if board[i][j] == 3 {
				target = []int{i, j}
			}
		}
	}

	dp := make([][]int, len(board))
	for i := range dp {
		sub := make([]int, len(board[0]))
		for j := range sub {
			sub[j] = 1<<31 - 1
		}
		dp[i] = sub
	}

	q := list.New()
	q.PushBack(start)
	dp[start[0]][start[1]] = 0

	for q.Len() > 0 {
		current := q.Front()
		q.Remove(current)
		cur := current.Value.([]int)

		for i := 1; i < len(dirs); i++ {
			row := cur[0] + dirs[i-1]
			col := cur[1] + dirs[i]

			if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
				continue
			}

			energy := dp[cur[0]][cur[1]]
			if board[row][col] == 1 {
				energy += c + 1
			} else {
				energy += 1
			}

			if dp[row][col] > energy {
				dp[row][col] = energy
				q.PushBack([]int{row, col})
			}
		}
	}

	return dp[target[0]][target[1]]
}
func solution4(arr [][]int) []int {
	var (
		dirs          = []int{0, 1, 0, -1, 0}
		area, maxArea = 0, 0
		max           = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		bfs = func(row, col int, arr [][]int) int {
			cnt := 1
			q := list.New()
			q.PushBack([]int{row, col})
			arr[row][col] = -1

			for q.Len() > 0 {
				current := q.Front()
				q.Remove(current)
				cur := current.Value.([]int)

				for i := 1; i < len(dirs); i++ {
					nextRow := cur[0] + dirs[i-1]
					nextCol := cur[1] + dirs[i]

					if nextRow < 0 || nextRow >= len(arr) || nextCol < 0 || nextCol >= len(arr[0]) {
						continue
					}

					if arr[nextRow][nextCol] == 0 {
						arr[nextRow][nextCol] = -1
						q.PushBack([]int{nextRow, nextCol})
						cnt++
					}
				}
			}
			return cnt
		}
	)

	for row := range arr {
		for col := range arr[row] {
			if arr[row][col] == 1 {
				maxArea = max(maxArea, bfs(row, col, arr))
				area++
			}
		}
	}

	return []int{area, maxArea}
}
func solution2(arr []int, k int, t int) int {
	var (
		answer = 0
		recur  = func(arr, out []int, cur, flag, depth int) {}
	)
	recur = func(arr []int, out []int, cur, flag, depth int) {
		if depth == len(out) {
			total := 0
			for _, v := range out {
				total += v
			}
			if total <= t {
				answer++
			}
			return
		} else {
			for i := cur; i < len(arr); i++ {
				if flag&(1<<i) == 0 {
					out[depth] = arr[i]
					recur(arr, out, i, flag|(1<<i), depth+1)
				}
			}
		}
	}

	for i := k; i <= len(arr); i++ {
		out := make([]int, i)
		recur(arr, out, 0, 0, 0)
	}

	return answer
}
func solution1(N int, sequence []int) int {
	var (
		answer = 0
		graph  = make([][]int, N+1)
		visit  = make([]bool, N+1)
		find   = func(from, to int) int {
			copy(visit, make([]bool, N+1))
			q := [][]int{{from, 0}}
			visit[from] = true

			for len(q) > 0 {
				cur := q[0]
				q = q[1:]
				if cur[0] == to {
					return cur[1]
				}
				for _, next := range graph[cur[0]] {
					if visit[next] == false {
						q = append(q, []int{next, cur[1] + 1})
					}
				}
			}
			return 0
		}
	)

	for i := 1; i <= N; i++ {
		graph[i] = []int{i - 1, i + 1}
	}
	graph[N][1] = 1
	graph[1][0] = N

	from := 1
	for i := 0; i < len(sequence); i++ {
		answer += find(from, sequence[i])
		from = sequence[i]
	}

	return answer
}
func solution(card []string, word []string) []string {
	var answer []string

	alpah := make([][]int, 0, 3)
	for _, v := range card {
		sub := make([]int, 26)
		for _, alphabet := range v {
			sub[alphabet-'A']++
		}
		alpah = append(alpah, sub)
	}
	answer = find(word, alpah, answer)
	return answer
}

func copyArr(src [][]int) [][]int {
	dest := make([][]int, len(src))
	for i := range src {
		dest[i] = make([]int, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}

func find(word []string, alpha [][]int, answer []string) []string {
Loop:
	for _, target := range word {
		visit := make([]bool, len(target))
		layer := make([]bool, 3)
		sub := copyArr(alpha)
		for i, alphabet := range sub {
			for j, v := range target {
				if visit[j] == false && alphabet[v-'A'] > 0 {
					alphabet[v-'A']--
					visit[j] = true
					if layer[i] == false {
						layer[i] = true
					}
				}
			}
		}
		for _, v := range visit {
			if v == false {
				continue Loop
			}
		}
		for _, v := range layer {
			if v == false {
				continue Loop
			}
		}
		answer = append(answer, target)
	}
	return answer
}
