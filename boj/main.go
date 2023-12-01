package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

/**

3
8
0 0
7 0
100
0 0
30 50
10
1 1
1 1


5
28
0
*/

func main() {
	solution7562()
}

func solution7562() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		total  int
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &total)
	for i := 0; i < total; i++ {
		size, cur, target := getInput(reader)
		if checkEqual(cur, target) {
			fmt.Fprintln(writer, 0)
			continue
		}
		minSteps := bfs7562(size, cur, target)
		fmt.Fprintln(writer, minSteps)
	}
}

func bfs7562(size int, cur, target []int) int {
	answer := 0
	dirs := [][]int{{-2, 1}, {-2, -1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	q := list.New()
	board[cur[0]][cur[1]] = -1
	q.PushBack(cur)
	for q.Len() > 0 {
		qSize := q.Len()
		for i := 0; i < qSize; i++ {
			c := q.Front()
			current := c.Value.([]int)
			q.Remove(c)
			for j := 0; j < len(dirs); j++ {
				row := current[0] + dirs[j][0]
				col := current[1] + dirs[j][1]

				if row < 0 || col < 0 || row >= size || col >= size || board[row][col] == -1 {
					continue
				}
				if checkEqual([]int{row, col}, target) {
					return answer + 1
				}
				board[row][col] = -1
				q.PushBack([]int{row, col})
			}
		}
		answer++
	}
	return answer
}
func checkEqual(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func getInput(reader *bufio.Reader) (int, []int, []int) {
	var (
		size, curX, curY, targetX, targetY int
	)
	fmt.Fscanln(reader, &size)
	fmt.Fscanln(reader, &curX, &curY)
	fmt.Fscanln(reader, &targetX, &targetY)

	return size, []int{curX, curY}, []int{targetX, targetY}
}
