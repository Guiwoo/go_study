package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	solution5014()
}

/**
F층 으로 이루어진 고층 건ㄹ물
G층 스타링크, S층 에 위치
버튼은 2개만 존재 위로 U만큼 이동 또는 D아래로
G층에 도달 못할시 use the stairs 출력
100,000 범위
*/

func solution5014() {
	var (
		reader          = bufio.NewReader(os.Stdin)
		writer          = bufio.NewWriter(os.Stdout)
		level, from, to int
		up, down        int
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &level, &from, &to, &up, &down)
	visit := make([]bool, level+1)
	q := list.New()
	q.PushBack(from)
	visit[from] = true
	cnt := 0
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			curVal := q.Front()
			q.Remove(curVal)
			cur := curVal.Value.(int)
			if cur == to {
				fmt.Fprintln(writer, cnt)
				return
			}
			for _, v := range []int{up, -down} {
				next := cur + v
				if next <= 0 || next > level || visit[next] {
					continue
				}
				if next == to {
					fmt.Fprintln(writer, cnt+1)
					return
				}
				q.PushBack(next)
				visit[next] = true
			}
		}
		cnt++
	}
	fmt.Fprintln(writer, "use the stairs")
}
