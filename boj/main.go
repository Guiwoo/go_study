package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cases int
	fmt.Fscanln(reader, &cases)

	for i := 0; i < cases; i++ {
		var row, col, lines int
		fmt.Fscanln(reader, &row, &col, &lines)

		arr := make([][]int, row)
		for r := range arr {
			arr[r] = make([]int, col)
		}

		for j := 0; j < lines; j++ {
			var tRow, tCol int
			fmt.Fscanln(reader, &tRow, &tCol)
			arr[tRow][tCol] = 1
		}

		writer.WriteString(bfs(arr) + "\n")
	}
}

func bfs(arr [][]int) string {
	dirs := []int{0, 1, 0, -1, 0}
	var ans int

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 1 {
				ans++
				arr[i][j] = -1
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]

						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr[i]) || arr[nRow][nCol] < 1 {
							continue
						}
						arr[nRow][nCol] = -1
						q = append(q, []int{nRow, nCol})
					}
				}
			}
		}
	}

	return strconv.Itoa(ans)
}
