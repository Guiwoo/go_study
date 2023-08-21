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
비숍은 좌우 대각선으로 움직일수 있다.
체스판이 주어지고 비숍을 놓을수 없는 위치가 존재한다면 서로가 서로를 잡을수 없는 위치에 놓을 수 있는 비숍의 최대 개수를 구하라
체스판 크기는 10이하이다.
비숍을 놓을수 있는 곳에는 1, 비숍을 놓을 수 없는 곳에는 0 이 빈칸을 사이에 두고 주어진다.

5
1 1 0 1 1
0 1 0 0 0
1 0 1 0 1
1 0 0 0 0
1 0 1 1 1

7
*/

func solution() {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	reader := bufio.NewReader(os.Stdin)
	var a int
	fmt.Fscanln(reader, &a)
	arr := make([][]int, a)
	diagonal := make([][][]int, a*2-1)
	for i := range arr {
		arr[i] = make([]int, a)
		for j := range arr[i] {
			var b int
			fmt.Fscan(reader, &b)
			arr[i][j] = b
			if b == 1 {
				diagonal[i+j] = append(diagonal[i+j], []int{i, j})
			}
		}
	}
	var answer int
	length := a * 2
	visit := make([]int, a*2)
	helper := func(depth, cnt int) {}
	helper = func(depth, cnt int) {
		if answer >= (cnt + length - depth) {
			return
		}
		if depth >= len(diagonal) {
			answer = max(answer, cnt)
			return
		} else {
			for _, v := range diagonal[depth] {
				target := (v[0] - v[1] + length) % length
				if visit[target] == 0 {
					visit[target] = 1
					helper(depth+2, cnt+1)
					visit[target] = 0
				}
			}
			helper(depth+2, cnt)
		}
	}
	helper(0, 0)
	ans := answer
	answer = 0
	helper(1, 0)
	fmt.Println(answer + ans)
}
