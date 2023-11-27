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
최대 100
*/

func solution() {
	var (
		reader   = bufio.NewReader(os.Stdin)
		row, col int
	)

	fmt.Fscanln(reader, &row, &col)

	maze := make([][]string, row)
	for i := 0; i < row; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		sub := make([]string, len(input))
		for i, v := range input {
			sub[i] = string(v)
		}
		maze[i] = sub
	}

	answer := bfs(0, 0, maze)
	fmt.Println(answer)
}

func bfs(row, col int, maze [][]string) int {
	type step struct {
		row, col, step int
	}

	dirs := []int{0, 1, 0, -1, 0}
	q := []step{step{}}
	maze[0][0] = "-1"

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.row == len(maze)-1 && cur.col == len(maze[0])-1 {
			return cur.step + 1
		}
		for i := 1; i < len(dirs); i++ {
			nRow := cur.row + dirs[i-1]
			nCol := cur.col + dirs[i]

			if nRow < 0 || nCol < 0 || nRow >= len(maze) || nCol >= len(maze[0]) ||
				maze[nRow][nCol] == "0" || maze[nRow][nCol] == "-1" {
				continue
			}
			maze[nRow][nCol] = "-1"
			q = append(q, step{nRow, nCol, cur.step + 1})
		}
	}
	return -1
}
