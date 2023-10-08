package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	solution()
}

/**
7
0110100
0110101
1110101
0000111
0100000
0111110
0111000
*/

func solution() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	arr := make([][]int, num)
	for i := range arr {
		arr[i] = make([]int, num)
		var input string
		fmt.Fscanln(reader, &input)
		for j := 0; j < len(input); j++ {
			arr[i][j] = int(input[j] - '0')
		}
	}

	var (
		ans    int
		amount []int
	)

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 1 {
				ans++
				arr[i][j] = 0
				amount = append(amount, bfs(i, j, arr))
			}
		}
	}

	fmt.Println(ans)
	sort.Ints(amount)
	for _, v := range amount {
		fmt.Println(v)
	}
}

func bfs(row, col int, arr [][]int) int {
	var ans int
	dirs := []int{0, 1, 0, -1, 0}

	q := [][]int{{row, col}}

	for len(q) > 0 {
		cur := q[0]
		ans++
		q = q[1:]
		for i := 1; i < len(dirs); i++ {
			nRow := cur[0] + dirs[i-1]
			nCol := cur[1] + dirs[i]

			if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr[0]) || arr[nRow][nCol] != 1 {
				continue
			}
			arr[nRow][nCol] = 0
			q = append(q, []int{nRow, nCol})
		}

	}

	return ans
}
