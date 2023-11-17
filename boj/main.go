package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solution()
}

//물을 붓는방법
// a->b
// a->c
// b->a
// b->c
// c->a
// c->b

// a물통 이 비어있으면 값을 추가
func solution() {
	var (
		reader = bufio.NewReader(os.Stdin)
		length = 3

		arr   = make([]int, 0, length)
		check = [201][201]bool{}

		answer = [201]bool{}
		dfs    = func(a, b, c int) {}
	)

	for i := 0; i < length; i++ {
		var a int
		fmt.Fscan(reader, &a)
		arr = append(arr, a)
	}

	dfs = func(a, b, c int) {
		if check[a][b] {
			return
		}
		if a == 0 {
			answer[c] = true
		}
		check[a][b] = true

		// a => b
		if a+b > arr[1] {
			dfs((a+b)-arr[1], arr[1], c)
		} else {
			dfs(0, a+b, c)
		}
		// b => a
		if a+b > arr[0] {
			dfs(arr[0], a+b-arr[0], c)
		} else {
			dfs(a+b, 0, c)
		}

		// c => a
		if a+c > arr[0] {
			dfs(arr[0], b, a+c-arr[0])
		} else {
			dfs(a+c, b, 0)
		}

		// c => b
		if c+b > arr[1] {
			dfs(a, arr[1], c+b-arr[1])
		} else {
			dfs(a, c+b, 0)
		}

		// a => c
		dfs(a, 0, b+c)
		// b => c
		dfs(0, b, a+c)
	}

	dfs(0, 0, arr[2])

	sb := strings.Builder{}
	for i := range answer {
		if answer[i] {
			sb.WriteString(fmt.Sprintf("%d ", i))
		}
	}
	fmt.Println(sb.String())
}
