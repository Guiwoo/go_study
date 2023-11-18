package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

/**
5
6 8 2 6 2
3 2 3 4 6
6 7 3 3 2
7 2 5 3 6
8 9 5 2 7
*/
// 5
/**
7
9 9 9 9 9 9 9
9 2 1 2 1 2 9
9 1 8 7 8 1 9
9 2 7 9 7 2 9
9 1 8 7 8 1 9
9 2 1 2 1 2 9
9 9 9 9 9 9 9
*/
// 6

func solution() {
	reader := bufio.NewReader(os.Stdin)
	var (
		n      int
		max    int
		answer = 0
	)
	fmt.Fscanln(reader, &n)

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		sub := make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(reader, &a)
			if a > max {
				max = a
			}
			sub[j] = a
		}
		arr[i] = sub
	}

	for i := max; i > 1; i-- {
		ans := findArea(arr, i)
		if ans > answer {
			answer = ans
		}
	}
	fmt.Println(answer)
}
func findArea(arr [][]int, rain int) int {
	var (
		length = len(arr)
		visit  = createVisit(length)
		dirs   = []int{0, 1, 0, -1, 0}
		answer = 0
	)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if arr[i][j] > rain && visit[i][j] == false {
				visit[i][j] = true
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						row := cur[0] + dirs[k-1]
						col := cur[1] + dirs[k]
						if row < 0 || row >= length || col < 0 || col >= length || visit[row][col] || arr[row][col] <= rain {
							continue
						}
						visit[row][col] = true
						q = append(q, []int{row, col})
					}
				}
				answer++
			}
		}
	}
	return answer
}

func createVisit(n int) [][]bool {
	visit := make([][]bool, n)
	for i := 0; i < n; i++ {
		visit[i] = make([]bool, n)
	}
	return visit
}
