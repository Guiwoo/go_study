package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution7569()
}

/**
첫 줄에는 상자의 크기를 나타내는 두 정수 M,N과 쌓아올려지는 상자의 수를 나타내는 H가 주어진다. M은 상자의 가로 칸의 수, N은 상자의 세로 칸의 수를 나타낸다.
단, 2 ≤ M ≤ 100, 2 ≤ N ≤ 100, 1 ≤ H ≤ 100 이다. 둘째 줄부터는 가장 밑의 상자부터 가장 위의 상자까지에 저장된 토마토들의 정보가 주어진다. 즉
, 둘째 줄부터 N개의 줄에는 하나의 상자에 담긴 토마토의 정보가 주어진다.
각 줄에는 상자 가로줄에 들어있는 토마토들의 상태가 M개의 정수로 주어진다. 정수 1은 익은 토마토, 정수 0 은 익지 않은 토마토, 정수 -1은 토마토가 들어있지 않은 칸을 나타낸다. 이러한 N개의 줄이 H번 반복하여 주어진다.

하나의 토마토에 인접한 곳은 위, 아래, 왼쪽, 오른쪽, 앞, 뒤 여섯 방향에 있는 토마토를 의미한다.

토마토가 하나 이상 있는 경우만 입력으로 주어진다.

여러분은 토마토가 모두 익을 때까지 최소 며칠이 걸리는지를 계산해서 출력해야 한다. 만약, 저장될 때부터 모든 토마토가 익어있는 상태이면 0을 출력해야 하고, 토마토가 모두 익지는 못하는 상황이면 -1을 출력해야 한다.

5 3 1
0 -1 0 0 0
-1 -1 0 1 1
0 0 0 1 1

-1

5 3 2
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 1 0 0
0 0 0 0 0

4

4 3 2
1 1 1 1
1 1 1 1
1 1 1 1
1 1 1 1
-1 -1 -1 -1
1 1 1 -1

0
*/

func solution7569() {
	var (
		reader              = bufio.NewReader(os.Stdin)
		writer              = bufio.NewWriter(os.Stdout)
		row, column, height int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &column, &row, &height)

	arr := make([][]int, row*height)

	dirRow := []int{0, 0, 1, -1, row, -row}
	dirCol := []int{1, -1, 0, 0, 0, 0}

	q := list.New()
	days := 0
	tomato := 0

	for i := range arr {
		sub := make([]int, column)
		line, _ := reader.ReadString('\n')
		lines := strings.Split(strings.TrimSpace(line), " ")
		for j, v := range lines {
			sub[j], _ = strconv.Atoi(v)
			if sub[j] == 1 {
				q.PushBack([]int{i, j})
			} else if sub[j] == 0 {
				tomato++
			}
		}
		arr[i] = sub
	}

	for q.Len() > 0 {
		size := q.Len()
		if tomato == 0 {
			fmt.Fprintln(writer, days)
			return
		}
		for i := 0; i < size; i++ {
			front := q.Front()
			q.Remove(front)
			cur := front.Value.([]int)
			for j := 0; j < len(dirCol); j++ {
				if j <= 3 {
					rRangeStart := (cur[0] / row) * row
					nRow := cur[0] + dirRow[j]
					nCol := cur[1] + dirCol[j]
					if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= column || arr[nRow][nCol] == -1 || arr[nRow][nCol] == 1 {
						continue
					}
					if rRangeStart <= nRow && nRow <= rRangeStart+row-1 {
						tomato--
						arr[nRow][nCol] = 1
						q.PushBack([]int{nRow, nCol})
					}
				} else {
					nRow := cur[0] + dirRow[j]
					nCol := cur[1] + dirCol[j]
					if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= column || arr[nRow][nCol] == -1 || arr[nRow][nCol] == 1 {
						continue
					}
					tomato--
					arr[nRow][nCol] = 1
					q.PushBack([]int{nRow, nCol})
				}
			}
		}
		days++
	}

	for i := range arr {
		for _, v := range arr[i] {
			if v == 0 {
				fmt.Fprintln(writer, -1)
				return
			}
		}
	}
	fmt.Fprintln(writer, days)
}
