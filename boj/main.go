package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var total int
	fmt.Fscanln(reader, &total)

	for i := 0; i < total; i++ {
		var r, c int
		fmt.Fscanln(reader, &r, &c)

		arr := make([][]string, r)

		for j := 0; j < r; j++ {
			var input string
			fmt.Fscanln(reader, &input)
			arr[j] = strings.Split(input, "")
		}

		writer.WriteString(bfs(arr) + "\n")
	}
}

func bfs(arr [][]string) string {
	ans := 0
	dirs := []int{0, 1, 0, -1, 0}

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == "#" {
				q := make([][]int, 1)
				q[0] = []int{i, j}

				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					arr[cur[0]][cur[1]] = "."
					//4방향 체크
					for k := 1; k < len(dirs); k++ {
						nextRow := cur[0] + dirs[k-1]
						nextCol := cur[1] + dirs[k]

						if nextRow < 0 || nextCol < 0 || nextRow >= len(arr) || nextCol >= len(arr[0]) {
							continue
						}
						if arr[nextRow][nextCol] == "#" {
							arr[nextRow][nextCol] = "."
							q = append(q, []int{nextRow, nextCol})
						}
					}
				}
				ans++
			}
		}
	}

	return strconv.Itoa(ans)
}
