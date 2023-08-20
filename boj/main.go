package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

/**
폴리오미노 크기가 1*1 인 정사각형을 여러 개 이어서 붙힌 도형 정사각형은 서로 겹치면 안 된다.
도형은 모두 연결되어 있어야 한다.
정사각형의 변끼리 연결되어 있어야 한다. 정사각형 4개를 이어붙힌 폴리오미노를 테리므노 라고 부른다. 5가지 유형이 있다
------
ㅁㅁㅁㅁ
------
ㅁㅁ
ㅁㅁ
------
ㅁ
ㅁ
ㅁㅁ
------
ㅁ
ㅁㅁ
 ㅁ
-------
ㅁㅁㅁ
 ㅁ
-------
n*m 종이 위에 테트로미노 하나를 놓으려고 한다. 종이는 1*1 크키 의 칸으로 이루어져있음 테르미노가 놓인 칸의 최대합을 구하세요
최대 사이즈 500*500 250k 여기서 5개로 검증한다. 1250k 연산수가 최대 1억 이 넘네 ?
시간 제한이 2초니깐 1억 연산을 2초안에 해야함 1초에 1억연산이 가능할까 ? 될꺼같은데
*/

/**
5 5
1 2 3 4 5
5 4 3 2 1
2 3 4 5 6
6 5 4 3 2
1 2 1 2 1

19

4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5

20

4 10
1 2 1 2 1 2 1 2 1 2
2 1 2 1 2 1 2 1 2 1
1 2 1 2 1 2 1 2 1 2
2 1 2 1 2 1 2 1 2 1

7
*/

func solution() {
	reader := bufio.NewReader(os.Stdin)
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	stick := func(arr [][]int, row, col int) (rs int) {
		//right 3steps more
		if col+3 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row][col+2]+arr[row][col+3])
		}
		//down 3steps more
		if row+3 < len(arr) {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+2][col]+arr[row+3][col])
		}
		return rs
	}
	square := func(arr [][]int, row, col int) (rs int) {
		// right 1step down 1step
		if row+1 < len(arr) && col+1 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row+1][col]+arr[row+1][col+1])
		}
		return rs
	}
	nieun := func(arr [][]int, row, col int) (rs int) {
		// right 2steps down 1step
		if row+1 < len(arr) && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+1][col+1]+arr[row+1][col+2])
		}
		// left 2steps down 1step
		if row+1 < len(arr) && col-2 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+1][col-1]+arr[row+1][col-2])
		}
		// top 1step right 2steps
		if row-1 >= 0 && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row-1][col]+arr[row-1][col+1]+arr[row-1][col+2])
		}
		// top 1step left 2steps
		if row-1 >= 0 && col-2 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row-1][col]+arr[row-1][col-1]+arr[row-1][col-2])
		}

		//right 1step down 2steps
		if row+2 < len(arr) && col+1 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row+1][col+1]+arr[row+2][col+1])
		}
		//right 1step top 2steps
		if row-2 >= 0 && col+1 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row-1][col+1]+arr[row-2][col+1])
		}
		// left 1step down 2steps
		if row+2 < len(arr) && col-1 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row][col-1]+arr[row+1][col-1]+arr[row+2][col-1])
		}
		// left 1step top 2steps
		if row-2 >= 0 && col-1 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row][col-1]+arr[row-1][col-1]+arr[row-2][col-1])
		}
		return rs
	}
	triangle := func(arr [][]int, row, col int) (rs int) {
		// right 2steps up 1step
		if row-1 >= 0 && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row][col+2]+arr[row-1][col+1])
		}
		//right 2steps down 1step
		if row+1 < len(arr) && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row][col+2]+arr[row+1][col+1])
		}

		// down 2steps right 1step
		if row+2 < len(arr) && col+1 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+2][col]+arr[row+1][col+1])
		}
		// down 2steps left 1step
		if row+2 < len(arr) && col-1 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+2][col]+arr[row+1][col-1])
		}
		return rs
	}
	etc := func(arr [][]int, row, col int) (rs int) {
		// right 2steps up 1step
		if row-1 >= 0 && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row-1][col+1]+arr[row-1][col+2])
		}
		//right 2steps down 1step
		if row+1 < len(arr) && col+2 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row][col+1]+arr[row+1][col+1]+arr[row+1][col+2])
		}

		//down 2steps right 1step
		if row+2 < len(arr) && col+1 < len(arr[0]) {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+1][col+1]+arr[row+2][col+1])
		}
		//down2steps left 1step
		if row+2 < len(arr) && col-1 >= 0 {
			rs = getMax(rs, arr[row][col]+arr[row+1][col]+arr[row+1][col-1]+arr[row+2][col-1])
		}
		return rs
	}

	check := func(arr [][]int, row, col int) int {
		result := getMax(square(arr, row, col), stick(arr, row, col))
		result2 := getMax(nieun(arr, row, col), triangle(arr, row, col))
		return getMax(result, getMax(result2, etc(arr, row, col)))
	}

	var answer int
	helper := func(arr [][]int, size int) {}
	helper = func(arr [][]int, size int) {
		if size == len(arr[0])*len(arr) {
			return
		}
		row, col := size/len(arr[0]), size%len(arr[0])
		answer = getMax(answer, check(arr, row, col))
		helper(arr, size+1)
	}

	var n, m int

	fmt.Fscanln(reader, &n, &m)

	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
		for j := range arr[i] {
			fmt.Fscan(reader, &arr[i][j])
		}
	}

	helper(arr, 0)
	fmt.Println(answer)
}
