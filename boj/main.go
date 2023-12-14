package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	boj14442()
}

func boj14442() {
	type step struct {
		x, y, steps, boom int
	}

	var (
		reader         = bufio.NewReader(os.Stdin)
		writer         = bufio.NewWriter(os.Stdout)
		row, col, boom int
		dirs           = []int{0, 1, 0, -1, 0}
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &row, &col, &boom)

	arr := make([][]int, row)
	visit := make([][][]bool, row)
	for i := range visit {
		sub := make([][]bool, col)
		for j := range sub {
			sub[j] = make([]bool, boom+1)
		}
		visit[i] = sub
	}

	for i := range arr {
		var (
			a   string
			sub = make([]int, col)
		)
		fmt.Fscanln(reader, &a)
		for j, v := range a {
			sub[j], _ = strconv.Atoi(string(v))
		}
		arr[i] = sub
	}

	q := list.New()
	q.PushBack(step{0, 0, 1, boom})
	for q.Len() > 0 {
		// 탈출조건 선정
		current := q.Front()
		q.Remove(current)
		cur := current.Value.(step)
		if cur.x == row-1 && cur.y == col-1 {
			fmt.Fprintln(writer, cur.steps)
			return
		}
		for i := 1; i < len(dirs); i++ {
			nRow := cur.x + dirs[i-1]
			nCol := cur.y + dirs[i]

			if nRow < 0 || nCol < 0 || nRow >= row || nCol >= col {
				continue
			}

			if nRow == row-1 && nCol == col-1 {
				fmt.Fprintln(writer, cur.steps+1)
				return
			}

			if arr[nRow][nCol] == 1 && cur.boom > 0 && visit[nRow][nCol][cur.boom-1] == false {
				visit[nRow][nCol][cur.boom-1] = true
				q.PushBack(step{nRow, nCol, cur.steps + 1, cur.boom - 1})
			} else if arr[nRow][nCol] == 0 {
				if cur.boom == 0 && visit[nRow][nCol][cur.boom] == false {
					visit[nRow][nCol][cur.boom] = true
					q.PushBack(step{nRow, nCol, cur.steps + 1, cur.boom})
				} else if cur.boom > 0 && visit[nRow][nCol][cur.boom] == false {
					visit[nRow][nCol][cur.boom] = true
					q.PushBack(step{nRow, nCol, cur.steps + 1, cur.boom})
				}
			}
		}
	}
	fmt.Fprintln(writer, -1)
}
