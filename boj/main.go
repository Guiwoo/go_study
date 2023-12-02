package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	solution15558()
}

func solution15558() {
	var (
		reader       = bufio.NewReader(os.Stdin)
		writer       = bufio.NewWriter(os.Stdout)
		length, jump int
	)

	fmt.Fscanln(reader, &length, &jump)
	defer writer.Flush()
	arr := setMap(reader, length)
	visit := make([][]bool, 2)
	for i := range visit {
		visit[i] = make([]bool, length)
	}

	q := list.New()
	q.PushBack([]int{0, 0}) // 0 left 1 right
	visit[0][0] = true
	sec := 0
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			el := q.Front()
			q.Remove(el)
			cur := el.Value.([]int)
			if cur[1] >= length {
				fmt.Fprintln(writer, 1)
				return
			}
			for k, v := range []int{1, -1, jump} {
				nextIdx := cur[1] + v
				if nextIdx >= length {
					fmt.Fprintln(writer, 1)
					return
				}
				if k == 2 {
					//건너뛴경우
					if cur[0] == 0 && arr[1][nextIdx] == 1 && !visit[1][nextIdx] {
						visit[1][nextIdx] = true
						q.PushBack([]int{1, nextIdx})
					} else if cur[0] == 1 && arr[0][nextIdx] == 1 && !visit[0][nextIdx] {
						visit[0][nextIdx] = true
						q.PushBack([]int{0, nextIdx})
					}
				} else {
					if nextIdx > sec && arr[cur[0]][nextIdx] == 1 && !visit[cur[0]][nextIdx] {
						visit[cur[0]][nextIdx] = true
						q.PushBack([]int{cur[0], nextIdx})
					}
				}
			}
		}
		sec++
	}
	fmt.Fprintln(writer, 0)
	return
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
