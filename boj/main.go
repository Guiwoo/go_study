package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

// v,o 를 loop 탐색으로 찾고 bfs v,o 숫자 비교하기
// o 가 더많다면 ? v를 쫗아내고 생존가능

/**
6 6
...#..
.##v#.
#v.#.#
#.o#.#
.###.#
...###

0 2

8 8
.######.
#..o...#
#.####.#
#.#v.#.#
#.#.o#o#
#o.##..#
#.v..v.#
.######.

3 1

9 12
.###.#####..
#.oo#...#v#.
#..o#.#.#.#.
#..##o#...#.
#.#v#o###.#.
#..#v#....#.
#...v#v####.
.####.#vv.o#
.......####.

3 5
*/

func solution() {
	var (
		row, col int
		reader   = bufio.NewReader(os.Stdin)
	)

	fmt.Fscanln(reader, &row, &col)

	area := setArea(row, col, reader)

	aliveSheep, aliveWolf := aliveCheck(area)
	fmt.Println(aliveSheep, aliveWolf)
}

func aliveCheck(area [][]rune) (int, int) {
	var totalSheep, totalWolf int
	for i := range area {
		for j := range area[i] {
			if area[i][j] == 'o' || area[i][j] == 'v' {
				sheep, wolf := bfs(area, i, j)

				totalSheep += sheep
				totalWolf += wolf
			}
		}
	}
	return totalSheep, totalWolf
}

func bfs(area [][]rune, row, col int) (int, int) {
	var (
		sheep, wolves int
	)
	if area[row][col] == 'o' {
		sheep++
	} else if area[row][col] == 'v' {
		wolves++
	}

	dirs := []int{0, 1, 0, -1, 0}
	q := [][]int{[]int{row, col}}
	area[row][col] = '#'
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 1; i < len(dirs); i++ {
			nRow := cur[0] + dirs[i-1]
			nCol := cur[1] + dirs[i]

			if nRow < 0 || nCol < 0 || nRow >= len(area) || nCol >= len(area[0]) || area[nRow][nCol] == '#' {
				continue
			}
			if area[nRow][nCol] == 'o' {
				sheep++
			} else if area[nRow][nCol] == 'v' {
				wolves++
			}
			area[nRow][nCol] = '#'
			q = append(q, []int{nRow, nCol})
		}
	}

	if sheep > wolves {
		return sheep, 0
	} else {
		return 0, wolves
	}
}

func setArea(row int, col int, reader *bufio.Reader) [][]rune {
	area := make([][]rune, row)
	for i := range area {
		var input string
		subArea := make([]rune, 0, col)
		fmt.Fscanln(reader, &input)
		for _, char := range input {
			subArea = append(subArea, char)
		}
		area[i] = subArea
	}
	return area
}
