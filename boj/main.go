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
	boj16933()
}

/**
N×M의 행렬로 표현되는 맵이 있다. 맵에서 0은 이동할 수 있는 곳을 나타내고, 1은 이동할 수 없는 벽이 있는 곳을 나타낸다. 당신은 (1, 1)에서 (N, M)의 위치까지 이동하려 하는데, 이때 최단 경로로 이동하려 한다. 최단경로는 맵에서 가장 적은 개수의 칸을 지나는 경로를 말하는데, 이때 시작하는 칸과 끝나는 칸도 포함해서 센다. 이동하지 않고 같은 칸에 머물러있는 경우도 가능하다. 이 경우도 방문한 칸의 개수가 하나 늘어나는 것으로 생각해야 한다.

이번 문제에서는 낮과 밤이 번갈아가면서 등장한다. 가장 처음에 이동할 때는 낮이고, 한 번 이동할 때마다 낮과 밤이 바뀌게 된다. 이동하지 않고 같은 칸에 머무르는 경우에도 낮과 밤이 바뀌게 된다.

만약에 이동하는 도중에 벽을 부수고 이동하는 것이 좀 더 경로가 짧아진다면, 벽을 K개 까지 부수고 이동하여도 된다. 단, 벽은 낮에만 부술 수 있다.

한 칸에서 이동할 수 있는 칸은 상하좌우로 인접한 칸이다.

맵이 주어졌을 때, 최단 경로를 구해 내는 프로그램을 작성하시오.
*/
/**
1 4 1
0010

5

1 4 1
0100

4

6 4 1
0100
1110
1000
0000
0111
0000

15

6 4 2
0100
1110
1000
0000
0111
0000

9
*/

func boj16933() {
	var (
		reader          = bufio.NewReader(os.Stdin)
		writer          = bufio.NewWriter(os.Stdout)
		row, col, booms int
		dirs            = []int{0, 1, 0, -1, 0}
		visit           [1001][1001][11]bool
	)
	type mover struct {
		x, y, steps, booms int
		isDay              bool
	}
	defer writer.Flush()

	fmt.Fscanln(reader, &row, &col, &booms)

	arr := make([][]int, row)
	for i := range arr {
		sub := make([]int, col)
		str, _ := reader.ReadString('\n')
		for j := range strings.TrimSpace(str) {
			sub[j], _ = strconv.Atoi(string(str[j]))
		}
		arr[i] = sub
	}

	q := list.New()
	q.PushBack(mover{0, 0, 1, booms, true})
	visit[0][0][booms] = true

	for q.Len() > 0 {
		current := q.Front()
		q.Remove(current)
		cur := current.Value.(mover)
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

			if arr[nRow][nCol] == 1 && cur.booms > 0 && visit[nRow][nCol][cur.booms-1] == false {
				if cur.isDay {
					visit[nRow][nCol][cur.booms-1] = true
					q.PushBack(mover{nRow, nCol, cur.steps + 1, cur.booms - 1, !cur.isDay})
				} else {
					q.PushBack(mover{cur.x, cur.y, cur.steps + 1, cur.booms, !cur.isDay})
				}
			}

			if arr[nRow][nCol] == 0 && visit[nRow][nCol][cur.booms] == false {
				visit[nRow][nCol][cur.booms] = true
				q.PushBack(mover{nRow, nCol, cur.steps + 1, cur.booms, !cur.isDay})
			}
		}
	}

	fmt.Fprintln(writer, -1)
}
