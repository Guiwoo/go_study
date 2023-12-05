package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	solution16948()
}

/**
r, c)라면, (r-2, c-1), (r-2, c+1), (r, c-2), (r, c+2), (r+2, c-1), (r+2, c+1)로
첫째 줄에 데스 나이트가 (r1, c1)에서 (r2, c2)로 이동하는 최소 이동 횟수를 출력한다. 이동할 수 없는 경우에는 -1을 출력한다.

7
6 6 0 1

4

6
5 1 0 5

-1

7
0 3 4 3

2
*/

func solution16948() {
	var (
		reader               = bufio.NewReader(os.Stdin)
		writer               = bufio.NewWriter(os.Stdout)
		size, steps          int
		startRow, startCol   int
		targetRow, targetCol int
		dirs                 = [][]int{{-2, -1}, {-2, +1}, {0, -2}, {0, 2}, {2, -1}, {2, 1}}
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &size)

	arr := make([][]int, size)

	for i := range arr {
		arr[i] = make([]int, size)
	}

	fmt.Fscanln(reader, &startRow, &startCol, &targetRow, &targetCol)

	q := list.New()
	arr[startRow][startCol] = -1
	q.PushBack([]int{startRow, startCol})

	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			value := q.Front()
			q.Remove(value)
			cur := value.Value.([]int)

			if cur[0] == targetRow && cur[1] == targetCol {
				fmt.Fprintln(writer, steps)
				return
			}

			for _, v := range dirs {
				nextRow := cur[0] + v[0]
				nextCol := cur[1] + v[1]

				if nextRow < 0 || nextCol < 0 || nextRow >= size || nextCol >= size || arr[nextRow][nextCol] == -1 {
					continue
				}

				arr[nextRow][nextCol] = -1
				if nextRow == targetRow && nextCol == targetCol {
					fmt.Fprintln(writer, steps+1)
					return
				}
				q.PushBack([]int{nextRow, nextCol})
			}
		}
		steps++
	}

	fmt.Fprintln(writer, -1)
}
