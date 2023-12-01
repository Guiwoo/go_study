package main

import "fmt"

// 격자칸을 칠해 가장 큰 정사각형을 만들려고 한다.
// k개이하

func main() {
	solution(3, []int{1, 2, 1, 3, 3, 2})
	solution(4, []int{3, 3, 2, 2, 1, 1, 4, 4})
}

func solution(n int, orders []int) {
	m := make(map[int]bool)
	arr := make([][]int, n)
	indexArr := make([]int, n+1)
	idx := 0

	for _, v := range orders {
		if ok, _ := m[v]; !ok {
			m[v] = true
			arr[idx] = []int{v, 2}
			indexArr[v] = idx
			idx++
		}
	}
	total := 0
	for _, v := range orders {
		index := indexArr[v]
		sub := 0
		if index == 0 {
			arr[index][1]--
			continue
		}
		arr[index][1]--
		for i := 0; i < index; i++ {
			if arr[i][1] > 0 {
				sub++
			}
		}
		total += sub
	}
	fmt.Println(total)
}
