package dfs_bfs

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func boj1260() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	arr := make([][]int, a+1)
	edge := 0
	for i := 0; i < b; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		arr[x] = append(arr[x], y)
		arr[y] = append(arr[y], x)
		edge++
	}

	for i := range arr {
		sort.Ints(arr[i])
	}
	var (
		dfs_answer []int
		bfs_answer []int
	)

	dfs := func(cur int, visit []bool) {}
	dfs = func(cur int, visit []bool) {
		target := arr[cur]
		for i := 0; i < len(target); i++ {
			if !visit[target[i]] {
				visit[target[i]] = true
				dfs_answer = append(dfs_answer, target[i])
				dfs(target[i], visit)
			}
		}
	}

	bfs := func(start int) {
		q := []int{start}
		visit := make([]bool, a+1)
		for len(q) > 0 {
			target := q[0]
			q = q[1:]
			if visit[target] {
				continue
			}
			visit[target] = true
			bfs_answer = append(bfs_answer, target)
			for i := 0; i < len(arr[target]); i++ {
				next := arr[target][i]
				if !visit[next] {
					q = append(q, next)
				}
			}
		}
	}
	visit := make([]bool, a+1)
	visit[c] = true

	dfs_answer = append(dfs_answer, c)
	dfs(c, visit)
	bfs(c)

	answer := strings.Builder{}
	for _, v := range dfs_answer {
		answer.WriteString(strconv.Itoa(v) + " ")
	}
	answer.WriteString("\n")
	for _, v := range bfs_answer {
		answer.WriteString(strconv.Itoa(v) + " ")
	}

	fmt.Println(answer.String())
}

func boj11724() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	arr := make([][]int, a+1)

	for i := 0; i < b; i++ {
		var start, end int
		fmt.Fscanln(reader, &start, &end)
		arr[start] = append(arr[start], end)
		arr[end] = append(arr[end], start)
	}
	visit := make([]bool, a+1)
	ans := 0
	for i := 1; i <= a; i++ {
		if !visit[i] {
			visit[i] = true
			boj11724Bfs(i, arr, visit)
			ans++
		}
	}
	fmt.Println(ans)
}

// dfs 가 조금더 빠른것으로 보임 왜그런지 확인하고 코드 작성해보기
func boj11724Bfs(start int, arr [][]int, visit []bool) {
	q := []int{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for i := 0; i < len(arr[cur]); i++ {
			if !visit[arr[cur][i]] {
				visit[arr[cur][i]] = true
				check := false
				for _, v := range arr[arr[cur][i]] {
					if !visit[v] {
						check = true
						break
					}
				}
				if check {
					q = append(q, arr[cur][i])
				}
			}

		}
	}
}

func boj11123() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var total int
	fmt.Fscanln(reader, &total)

	for i := 0; i < total; i++ {
		var r, c int
		fmt.Fscanln(reader, &r, &c)

		arr := make([][]string, r)

		for j := 0; j < r; j++ {
			var input string
			fmt.Fscanln(reader, &input)
			arr[j] = strings.Split(input, "")
		}

		writer.WriteString(boj11123Bfs(arr) + "\n")
	}
}

func boj11123Bfs(arr [][]string) string {
	ans := 0
	dirs := []int{0, 1, 0, -1, 0}

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == "#" {
				q := make([][]int, 1)
				q[0] = []int{i, j}

				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					arr[cur[0]][cur[1]] = "."
					//4방향 체크
					for k := 1; k < len(dirs); k++ {
						nextRow := cur[0] + dirs[k-1]
						nextCol := cur[1] + dirs[k]

						if nextRow < 0 || nextCol < 0 || nextRow >= len(arr) || nextCol >= len(arr[0]) {
							continue
						}
						if arr[nextRow][nextCol] == "#" {
							arr[nextRow][nextCol] = "."
							q = append(q, []int{nextRow, nextCol})
						}
					}
				}
				ans++
			}
		}
	}

	return strconv.Itoa(ans)
}

func boj11123Dfs(arr [][]string) string {
	ans := 0
	recur := func(i, j int) {}

	recur = func(i, j int) {
		if i < 0 || j < 0 || i >= len(arr) || j >= len(arr[0]) || arr[i][j] != "#" {
			return
		}
		if arr[i][j] == "#" {
			arr[i][j] = "."
		}
		recur(i-1, j)
		recur(i, j-1)
		recur(i+1, j)
		recur(i, j+1)
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "#" {
				recur(i, j)
				ans++
			}
		}
	}
	return strconv.Itoa(ans)
}

func boj1012() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cases int
	fmt.Fscanln(reader, &cases)

	for i := 0; i < cases; i++ {
		var row, col, lines int
		fmt.Fscanln(reader, &row, &col, &lines)

		arr := make([][]int, row)
		for r := range arr {
			arr[r] = make([]int, col)
		}

		for j := 0; j < lines; j++ {
			var tRow, tCol int
			fmt.Fscanln(reader, &tRow, &tCol)
			arr[tRow][tCol] = 1
		}

		writer.WriteString(boj1012Bfs(arr) + "\n")
	}
}

func boj1012Bfs(arr [][]int) string {
	dirs := []int{0, 1, 0, -1, 0}
	var ans int

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 1 {
				ans++
				arr[i][j] = -1
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]

						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr[i]) || arr[nRow][nCol] < 1 {
							continue
						}
						arr[nRow][nCol] = -1
						q = append(q, []int{nRow, nCol})
					}
				}
			}
		}
	}

	return strconv.Itoa(ans)
}

func boj2667() {
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
				amount = append(amount, bfs2667(i, j, arr))
			}
		}
	}

	fmt.Println(ans)
	sort.Ints(amount)
	for _, v := range amount {
		fmt.Println(v)
	}
}

func bfs2667(row, col int, arr [][]int) int {
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

func boj2583() {
	reader := bufio.NewReader(os.Stdin)

	var size int
	fmt.Fscanln(reader, &size)

	graph := make([][]int, size)

	arr := make([][]int, size)
	for i := range arr {
		sub := make([]int, size)
		for j := range sub {
			var a int
			fmt.Fscan(reader, &a)
			if a == 1 {
				graph[i] = append(graph[i], j)
			}
			sub[j] = a
		}
		arr[i] = sub
	}

	answer := make([][]int, size)

	for i := 0; i < len(graph); i++ {
		visit := make([]bool, len(graph))
		boj2583Dfs(i, graph, visit)
		sub := make([]int, size)
		for j := 0; j < len(sub); j++ {
			if visit[j] {
				sub[j] = 1
			}
		}
		answer[i] = sub
	}

	for _, v := range answer {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Println()
	}
}

func boj2583Dfs(idx int, graph [][]int, visit []bool) {
	for i := 0; i < len(graph[idx]); i++ {
		if !visit[graph[idx][i]] {
			visit[graph[idx][i]] = true
			boj2583Dfs(graph[idx][i], graph, visit)
		}
	}
}

func boj10026() {
	reader := bufio.NewReader(os.Stdin)

	var len int
	fmt.Fscanln(reader, &len)

	arr := make([][]string, len)

	for i := range arr {
		sub := make([]string, len)
		var a string
		fmt.Fscanln(reader, &a)
		for i, v := range a {
			sub[i] = string(v)
		}
		arr[i] = sub
	}

	abNormal := bfsAbnormal(copySlice(arr))
	normal := bfsNormal(copySlice(arr))

	fmt.Println(normal, abNormal)
}

func copySlice(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i, sub := range src {
		dst[i] = make([]string, len(sub))
		copy(dst[i], sub)
	}
	return dst
}

func bfsNormal(arr [][]string) int {
	var (
		answer int
		dirs   []int = []int{0, 1, 0, -1, 0}
	)
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != "A" {
				// do something
				first := arr[i][j]
				arr[i][j] = "A"
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]
						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr) || arr[nRow][nCol] == "A" || arr[nRow][nCol] != first {
							continue
						}
						arr[nRow][nCol] = "A"
						q = append(q, []int{nRow, nCol})
					}
				}

				answer++
			}
		}
	}
	return answer
}

func bfsAbnormal(arr [][]string) int {
	var (
		answer int
		dirs   []int = []int{0, 1, 0, -1, 0}
	)
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != "A" {
				// do something
				first := arr[i][j]
				arr[i][j] = "A"
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						nRow := cur[0] + dirs[k-1]
						nCol := cur[1] + dirs[k]
						if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr) || arr[nRow][nCol] == "A" {
							continue
						}
						if first == "B" && arr[nRow][nCol] != first {
							continue
						}
						if first != "B" && arr[nRow][nCol] == "B" {
							continue
						}
						arr[nRow][nCol] = "A"
						q = append(q, []int{nRow, nCol})
					}
				}

				answer++
			}
		}
	}
	return answer
}

func boj_2251() {
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

func solution2468() {
	reader := bufio.NewReader(os.Stdin)
	var (
		n      int
		max    int
		answer = 0
	)
	fmt.Fscanln(reader, &n)

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		sub := make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(reader, &a)
			if a > max {
				max = a
			}
			sub[j] = a
		}
		arr[i] = sub
	}

	for i := max; i > 1; i-- {
		ans := findArea(arr, i)
		if ans > answer {
			answer = ans
		}
	}
	fmt.Println(answer)
}
func findArea(arr [][]int, rain int) int {
	var (
		length = len(arr)
		visit  = createVisit(length)
		dirs   = []int{0, 1, 0, -1, 0}
		answer = 0
	)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if arr[i][j] > rain && visit[i][j] == false {
				visit[i][j] = true
				q := [][]int{{i, j}}
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for k := 1; k < len(dirs); k++ {
						row := cur[0] + dirs[k-1]
						col := cur[1] + dirs[k]
						if row < 0 || row >= length || col < 0 || col >= length || visit[row][col] || arr[row][col] <= rain {
							continue
						}
						visit[row][col] = true
						q = append(q, []int{row, col})
					}
				}
				answer++
			}
		}
	}
	return answer
}

func createVisit(n int) [][]bool {
	visit := make([][]bool, n)
	for i := 0; i < n; i++ {
		visit[i] = make([]bool, n)
	}
	return visit
}

func solution3184() {
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

func solution14615() {
	var (
		areas, edges, input int
		reader              = bufio.NewReader(os.Stdin)
		ans                 = strings.Builder{}
		visit               = [2][100001]bool{}
	)

	fmt.Fscanln(reader, &areas, &edges)
	graph, reverse := setGraph(areas, edges, reader)

	findWay(1, areas, 0, graph, visit)
	findWay(areas, 1, 1, reverse, visit)

	fmt.Fscanln(reader, &input)
	for i := 0; i < input; i++ {
		var bomb int
		fmt.Fscanln(reader, &bomb)
		if visit[0][bomb] && visit[1][bomb] {
			ans.WriteString("Defend the CTP\n")
		} else {
			ans.WriteString("Destroyed the CTP\n")
		}
	}
	fmt.Println(ans.String())
}
func findWay(from, to, way int, graph map[int][]int, visit [2][100001]bool) {
	visit[way][from] = true
	q := []int{from}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < len(graph[cur]); i++ {
			next := graph[cur][i]
			if visit[way][next] == false {
				visit[way][next] = true
				q = append(q, next)
			}
		}
	}
	return
}

func setGraph(areas, edges int, reader *bufio.Reader) (map[int][]int, map[int][]int) {
	graph := make(map[int][]int)
	reverse := make(map[int][]int)
	for i := 0; i <= areas; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Fscanln(reader, &from, &to)
		graph[from] = append(graph[from], to)
		reverse[to] = append(reverse[to], from)
	}
	return graph, reverse
}

func solution1389() {
	var (
		reader           = bufio.NewReader(os.Stdin)
		users, relations int
		index            int
	)

	fmt.Fscanln(reader, &users, &relations)
	graph := make([][]int, users+1)
	answer := make([]int, users+1)
	for i := range answer {
		answer[i] = math.MaxInt
	}

	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < relations; i++ {
		var who, know int
		fmt.Fscanln(reader, &who, &know)
		graph[who] = append(graph[who], know)
		graph[know] = append(graph[know], who)
	}

	for user := 1; user <= users; user++ {
		var (
			cnt   = 0
			q     = []int{user}
			visit = make([]bool, users+1)
			run   = 0
		)

		visit[user] = true

		for len(q) > 0 {
			size := len(q)
			for i := 0; i < size; i++ {
				cur := q[0]
				cnt += run
				q = q[1:]
				for _, v := range graph[cur] {
					if visit[v] == false {
						visit[v] = true
						q = append(q, v)
					}
				}
			}
			run++
		}

		if answer[index] > cnt {
			index = user
			answer[index] = cnt
		}
	}

	fmt.Println(index)
}

func solution2178() {
	var (
		reader   = bufio.NewReader(os.Stdin)
		row, col int
	)

	fmt.Fscanln(reader, &row, &col)

	maze := make([][]string, row)
	for i := 0; i < row; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		sub := make([]string, len(input))
		for i, v := range input {
			sub[i] = string(v)
		}
		maze[i] = sub
	}

	answer := bfs2178(0, 0, maze)
	fmt.Println(answer)
}

func bfs2178(row, col int, maze [][]string) int {
	type step struct {
		row, col, step int
	}

	dirs := []int{0, 1, 0, -1, 0}
	q := []step{step{}}
	maze[0][0] = "-1"

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.row == len(maze)-1 && cur.col == len(maze[0])-1 {
			return cur.step + 1
		}
		for i := 1; i < len(dirs); i++ {
			nRow := cur.row + dirs[i-1]
			nCol := cur.col + dirs[i]

			if nRow < 0 || nCol < 0 || nRow >= len(maze) || nCol >= len(maze[0]) ||
				maze[nRow][nCol] == "0" || maze[nRow][nCol] == "-1" {
				continue
			}
			maze[nRow][nCol] = "-1"
			q = append(q, step{nRow, nCol, cur.step + 1})
		}
	}
	return -1
}

func solution7556() {
	var (
		reader                   = bufio.NewReader(os.Stdin)
		row, col, tomatoes, days int
		q                        = make([][]int, 0)
	)
	fmt.Fscanln(reader, &col, &row)

	arr := make([][]string, row)
	for i := range arr {
		text, _, _ := reader.ReadLine()
		sub := strings.Split(string(text), " ")
		for j, v := range sub {
			if v == "0" {
				tomatoes++
			} else if v == "1" {
				q = append(q, []int{i, j})
			}
		}
		arr[i] = sub
	}

	for len(q) > 0 {
		if tomatoes == 0 {
			break
		}
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			bfs7556(&q, arr, cur[0]-1, cur[1], &tomatoes)
			bfs7556(&q, arr, cur[0], cur[1]-1, &tomatoes)
			bfs7556(&q, arr, cur[0]+1, cur[1], &tomatoes)
			bfs7556(&q, arr, cur[0], cur[1]+1, &tomatoes)
		}
		days++
	}

	if tomatoes == 0 {
		fmt.Println(days)
	} else {
		fmt.Println(-1)
	}
}

func bfs7556(q *[][]int, arr [][]string, nRow, nCol int, tomatoes *int) {
	if nRow < 0 || nCol < 0 || nRow >= len(arr) || nCol >= len(arr[0]) {
		return
	}
	if arr[nRow][nCol] == "0" {
		arr[nRow][nCol] = "-99"
		*tomatoes--
		*q = append(*q, []int{nRow, nCol})
	}
}

func solution1679() {
	var (
		reader        = bufio.NewReader(os.Stdin)
		mover, target int
	)

	fmt.Fscanln(reader, &mover, &target)
	bfs1679(mover, target)
}

/*
*
수빈이가 5-10-9-18-17 순으로 가면 4초만에 동생을 찾을 수 있다.
4,6,10 1초

3,8
7,12 2초
*/
func bfs1679(move, target int) int {
	type mover struct {
		move, sec int
	}
	visit := make([]int, 100001)
	q := []int{move}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == target {
			fmt.Println(visit[target])
			break
		}
		for _, v := range []int{1, -1, cur} {
			n := v + cur
			if 0 <= n && n <= 100000 && visit[n] == 0 {
				visit[n] = visit[cur] + 1
				q = append(q, n)
			}
		}
	}
	return visit[target]
}

func solution12851() {
	var (
		reader        = bufio.NewReader(os.Stdin)
		mover, target int
	)

	fmt.Fscanln(reader, &mover, &target)
	bfs12851(mover, target)
}

/**
5 10 9 18 17
5 4 8 16 17

0 1 2 3
0 1 2 3
*/

func bfs12851(move, target int) {
	visit := make([]int, 100001)
	for i := range visit {
		visit[i] = -1
	}
	visit[move] = 0
	path := make([]int, 100001)
	path[move] = 1
	q := []int{move}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range []int{1, -1, cur} {
			n := v + cur
			if 0 <= n && n <= 100000 {
				if visit[n] == -1 {
					visit[n] = visit[cur] + 1
					path[n] = path[cur]
					q = append(q, n)
				} else {
					if visit[n] == visit[cur]+1 {
						path[n] += path[cur]
					}
				}
			}
		}
	}
	fmt.Println(visit[target], path[target])
}

func solution13549() {
	var (
		reader        = bufio.NewReader(os.Stdin)
		mover, target int
	)

	fmt.Fscanln(reader, &mover, &target)
	bfs13549(mover, target)
}

func bfs13549(move, target int) {
	visit := make([]int, 100001)
	for i := range visit {
		visit[i] = -1
	}
	visit[move] = 0
	q := []int{move}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == target {
			fmt.Println(visit[target])
			break
		}
		for _, v := range []int{cur, 1, -1} {
			n := v + cur
			if 0 <= n && n <= 100000 {
				if v == cur {
					visit[n] = visit[cur]
					q = append(q, n)
				} else {
					num := visit[cur] + 1
					if visit[n] == -1 {
						visit[n] = num
						q = append(q, n)
					} else if visit[n] > num {
						visit[n] = num
						q = append(q, n)
					}
				}
			}
		}
	}
}

var (
	reader   = bufio.NewReader(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	from, to int

	visit [100001]int
	path  [100001]int
)

func solution13913() {

	defer writer.Flush()
	fmt.Fscan(reader, &from, &to)

	q := list.New()

	visit[from] = 1
	path[from] = from
	q.PushBack(from)

	for q.Len() > 0 {
		cur := q.Front().Value.(int)
		q.Remove(q.Front())

		if cur == to {
			break
		}
		for _, v := range []int{-1, cur, 1} {
			next := v + cur
			if next >= 0 && next <= 100000 && visit[next] == 0 {
				visit[next] = visit[cur] + 1
				path[next] = cur
				q.PushBack(next)
			}
		}
	}

	fmt.Fprintln(writer, visit[to]-1)
	getPath(to)
}

func getPath(u int) {
	if u == path[u] {
		fmt.Fprint(writer, u, " ")
		return
	}
	getPath(path[u])
	fmt.Fprint(writer, u, " ")
}

func solution7562() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		total  int
	)
	defer writer.Flush()

	fmt.Fscanln(reader, &total)
	for i := 0; i < total; i++ {
		size, cur, target := getInput(reader)
		if checkEqual(cur, target) {
			fmt.Fprintln(writer, 0)
			continue
		}
		minSteps := bfs7562(size, cur, target)
		fmt.Fprintln(writer, minSteps)
	}
}

func bfs7562(size int, cur, target []int) int {
	answer := 0
	dirs := [][]int{{-2, 1}, {-2, -1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	q := list.New()
	board[cur[0]][cur[1]] = -1
	q.PushBack(cur)
	for q.Len() > 0 {
		qSize := q.Len()
		for i := 0; i < qSize; i++ {
			c := q.Front()
			current := c.Value.([]int)
			q.Remove(c)
			for j := 0; j < len(dirs); j++ {
				row := current[0] + dirs[j][0]
				col := current[1] + dirs[j][1]

				if row < 0 || col < 0 || row >= size || col >= size || board[row][col] == -1 {
					continue
				}
				if checkEqual([]int{row, col}, target) {
					return answer + 1
				}
				board[row][col] = -1
				q.PushBack([]int{row, col})
			}
		}
		answer++
	}
	return answer
}
func checkEqual(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func getInput(reader *bufio.Reader) (int, []int, []int) {
	var (
		size, curX, curY, targetX, targetY int
	)
	fmt.Fscanln(reader, &size)
	fmt.Fscanln(reader, &curX, &curY)
	fmt.Fscanln(reader, &targetX, &targetY)

	return size, []int{curX, curY}, []int{targetX, targetY}
}
