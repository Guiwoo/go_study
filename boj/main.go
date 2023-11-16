package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)

	var len int
	fmt.Fscanln(reader, &len)

	arr := make([][]string, len)

	for i := range arr {
		sub := make([]string, len)
		var a string
		fmt.Fscanln(reader, &a)
		for i, v := range a {
			sub[i] = string(v)
		}
		arr[i] = sub
	}

	abNormal := bfsAbnormal(copySlice(arr))
	normal := bfsNormal(copySlice(arr))

	fmt.Println(normal, abNormal)
}

func copySlice(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i, sub := range src {
		dst[i] = make([]string, len(sub))
		copy(dst[i], sub)
	}
	return dst
}

func bfsNormal(arr [][]string) int {
	var (
		answer int
		dirs   []int = []int{0, 1, 0, -1, 0}
	)
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != "A" {
				// do something
				first := arr[i][j]
				arr[i][j] = "A"
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]
						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr) || arr[nRow][nCol] == "A" || arr[nRow][nCol] != first {
							continue
						}
						arr[nRow][nCol] = "A"
						q = append(q, []int{nRow, nCol})
					}
				}

				answer++
			}
		}
	}
	return answer
}

func bfsAbnormal(arr [][]string) int {
	var (
		answer int
		dirs   []int = []int{0, 1, 0, -1, 0}
	)
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != "A" {
				// do something
				first := arr[i][j]
				arr[i][j] = "A"
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]
						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr) || arr[nRow][nCol] == "A" {
							continue
						}
						if first == "B" && arr[nRow][nCol] != first {
							continue
						}
						if first != "B" && arr[nRow][nCol] == "B" {
							continue
						}
						arr[nRow][nCol] = "A"
						q = append(q, []int{nRow, nCol})
					}
				}

				answer++
			}
		}
	}
	return answer
}
