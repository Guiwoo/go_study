package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	row, col int
}

func main() {
	solution15558()
}

func solution15558() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var length, jump int
	fmt.Fscanln(reader, &length, &jump)

	arr := setMap(reader, length)
	visit := make([][]bool, 2)
	for i := range visit {
		visit[i] = make([]bool, length)
	}

	q := list.New()
	q.PushBack(Position{0, 0}) // 0 left 1 right
	visit[0][0] = true

	sec := 0
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			cur := q.Remove(q.Front()).(Position)
			if cur.col >= length {
				fmt.Fprintln(writer, 1)
				return
			}
			for _, v := range []int{1, -1, jump} {
				nextIdx := cur.col + v
				if nextIdx >= length {
					fmt.Fprintln(writer, 1)
					return
				}
				if v == jump {
					// 건너뛴 경우
					otherRow := 1 - cur.row
					if arr[otherRow][nextIdx] == 1 && !visit[otherRow][nextIdx] {
						visit[otherRow][nextIdx] = true
						q.PushBack(Position{otherRow, nextIdx})
					}
				} else if nextIdx > sec && arr[cur.row][nextIdx] == 1 && !visit[cur.row][nextIdx] {
					visit[cur.row][nextIdx] = true
					q.PushBack(Position{cur.row, nextIdx})
				}
			}
		}
		sec++
	}
	fmt.Fprintln(writer, 0)
}

func setMap(reader *bufio.Reader, length int) [][]int {
	arr := make([][]int, 2)
	for k := range arr {
		sub := make([]int, length)
		var num string
		fmt.Fscanln(reader, &num)
		for i, v := range num {
			sub[i], _ = strconv.Atoi(string(v))
		}
		arr[k] = sub
	}
	return arr
}
