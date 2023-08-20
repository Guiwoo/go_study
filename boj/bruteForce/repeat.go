package bruteForce

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func repeat_2309() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	arr := make([]int, 9)
	var (
		total, x, y int
	)
	for i := range arr {
		var a int
		fmt.Fscanln(reader, &a)
		arr[i] = a
		total += a
	}

	sort.Ints(arr)
Loop:
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if total-(arr[i]+arr[j]) == 100 {
				x, y = i, j
				break Loop
			}
		}
	}
	fmt.Println(x, y)

	for i := range arr {
		if i != x && i != y {
			fmt.Println(arr[i])
		}
	}
}

func repeat_6603() {
	reader := bufio.NewReader(os.Stdin)
	answer := strings.Builder{}

	bitmask := func(arr, rs []string, flag, start int) {}
	bitmask = func(arr, rs []string, flag, start int) {
		if len(rs) >= 6 {
			ans := ""
			for _, v := range rs {
				ans += v + " "
			}
			answer.WriteString(ans + "\n")
			return
		} else {
			for i := start; i < len(arr); i++ {
				if flag&(1<<i) != 0 {
					continue
				}
				bitmask(arr, append(rs, arr[i]), flag|(1<<i), i+1)
			}
		}
	}

	for {
		line, _ := reader.ReadString('\n')
		if line[0] == '0' {
			break
		}
		input := strings.TrimSpace(line)
		target := strings.Split(input, " ")[1:]
		var rs []string
		bitmask(target, rs, 0, 0)
	}
	fmt.Println(answer.String())
}

func repeat_1182() {
	reader := bufio.NewReader(os.Stdin)

	var (
		n, answer, sum int
	)

	fmt.Fscanln(reader, &n, &sum)

	arr := make([]int, n)
	for i := range arr {
		var a int
		fmt.Fscan(reader, &a)
		arr[i] = a
	}

	find := func(arr []int, total, idx, flag int) {}
	find = func(arr []int, total, idx, flag int) {
		if idx >= len(arr) {
			if total == sum && flag > 0 {
				answer++
			}
			return
		} else {
			find(arr, total+arr[idx], idx+1, flag|(1<<idx))
			find(arr, total, idx+1, flag)
		}

	}
	find(arr, 0, 0, 0)
	fmt.Println(answer)
}

func repeat_2580() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([][]int, 9)
	for i := range arr {
		sub := make([]int, 9)
		for j := range sub {
			var a int
			fmt.Fscan(reader, &a)
			sub[j] = a
		}
		arr[i] = sub
	}
	isPossible := func(arr [][]int, num, row, col int) bool {
		//가로줄 체크
		for i := 0; i < 9; i++ {
			if arr[row][i] == num {
				return false
			}
		}

		//세로줄 체크
		for i := 0; i < 9; i++ {
			if arr[i][col] == num {
				return false
			}
		}

		// 3*3 체크
		row = (row / 3) * 3
		col = (col / 3) * 3
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				if arr[i][j] == num {
					return false
				}
			}
		}
		return true
	}

	dfs := func(arr [][]int, cal int) {}
	dfs = func(arr [][]int, cal int) {
		if cal >= 81 {
			for i := range arr {
				for j := range arr[i] {
					fmt.Printf("%d ", arr[i][j])
				}
				fmt.Println()
			}
			os.Exit(0)
		}
		row := cal / 9
		col := cal % 9
		if arr[row][col] != 0 {
			dfs(arr, cal+1)
		} else {
			for i := 1; i <= 9; i++ {
				if isPossible(arr, i, row, col) {
					arr[row][col] = i
					dfs(arr, cal+1)
					arr[row][col] = 0
				}
			}
		}
	}
	dfs(arr, 0)
}

func repeat_1251() {
	reader := bufio.NewReader(os.Stdin)
	reverse := func(s string) (ans string) {
		for i := len(s) - 1; i >= 0; i-- {
			ans += string(s[i])
		}
		return ans
	}
	strCompare := func(a string, b string, n []int) string {
		b1 := reverse(b[:n[0]])
		b2 := reverse(b[n[0]:n[1]])
		b3 := reverse(b[n[1]:])
		if a > (b1 + b2 + b3) {
			return b1 + b2 + b3
		}
		return a
	}

	var input string
	answer := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
	fmt.Fscanln(reader, &input)

	recur := func(n int, rs []int, start, flag int) {}
	recur = func(n int, rs []int, start, flag int) {
		if len(rs) == 2 {
			answer = strCompare(answer, input, rs)
			return
		} else {
			for i := start; i < n; i++ {
				if flag&(1<<i) == 0 {
					recur(n, append(rs, i), i+1, flag|(1<<i))
				}
			}
		}
	}
	recur(len(input), []int{}, 1, 0)
	fmt.Println(answer)
}

func repeat_2529() {
	reader := bufio.NewReader(os.Stdin)
	var a int
	fmt.Fscanln(reader, &a)
	arr := make([]string, a)
	for i := range arr {
		var b string
		fmt.Fscan(reader, &b)
		arr[i] = b
	}

	min, max := "9999999999", "0"

	recur := func(rs []string, depth, flag int) {}
	recur = func(rs []string, depth, flag int) {
		if len(rs) >= a+1 {
			result := strings.Join(rs, "")
			if min > result {
				min = result
			}
			if max < result {
				max = result
			}
			return
		} else {
			for i := 0; i <= 9; i++ {
				if depth == 0 {
					recur(append(rs, fmt.Sprintf("%d", i)), depth+1, flag|(1<<i))
				} else {
					if flag&(1<<i) == 0 {
						next := fmt.Sprintf("%d", i)
						switch arr[depth-1] {
						case "<":
							if rs[depth-1] < next {
								recur(append(rs, next), depth+1, flag|(1<<i))
							}
						default:
							if rs[depth-1] > next {
								recur(append(rs, next), depth+1, flag|(1<<i))
							}
						}
					}
				}
			}
		}
	}
	recur([]string{}, 0, 0)
	fmt.Printf("%s\n%s", max, min)
}

func repeat_14888() {
	reader := bufio.NewReader(os.Stdin)

	var a int
	fmt.Fscanln(reader, &a)

	arr := make([]int, a)
	for i := range arr {
		fmt.Fscan(reader, &a)
		arr[i] = a
	}
	oper := make([]int, 4)
	for i := range oper {
		fmt.Fscan(reader, &a)
		oper[i] = a
	}
	min, max := 1<<31, -(1 << 31)

	recur := func(idx, sum int) {}
	recur = func(idx, sum int) {
		if idx == len(arr) {
			if min > sum {
				min = sum
			}
			if max < sum {
				max = sum
			}
			return
		} else if idx == 0 {
			recur(idx+1, arr[idx])
		} else {
			for i := 0; i < len(oper); i++ {
				if oper[i] > 0 {
					switch i {
					case 0:
						oper[i]--
						recur(idx+1, sum+arr[idx])
						oper[i]++
					case 1:
						oper[i]--
						recur(idx+1, sum-arr[idx])
						oper[i]++
					case 2:
						oper[i]--
						recur(idx+1, sum*arr[idx])
						oper[i]++
					default:
						oper[i]--
						recur(idx+1, sum/arr[idx])
						oper[i]++
					}
				}
			}
		}
	}
	recur(0, 0)
	fmt.Printf("%d\n%d", max, min)
}

func repeat_14889() {
	reader := bufio.NewReader(os.Stdin)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var a int
	fmt.Fscanln(reader, &a)

	arr := make([][]int, a)
	row := make([]int, a)
	col := make([]int, a)
	total := 0
	for i := range arr {
		sub := make([]int, a)
		for j := range sub {
			var b int
			fmt.Fscan(reader, &b)
			sub[j] = b
			total += b
			row[i] += b
			col[j] += b
		}
		arr[i] = sub
	}
	fmt.Println(row, col, total)
	min := 1 << 31
	recur := func(start, depth, sum, flag int) {}
	recur = func(start, depth, sum, flag int) {
		if depth == a/2 {
			target := abs(total - sum)
			if target < min {
				min = target
			}
			return
		} else {
			for i := start; i < a; i++ {
				if flag&(1<<i) == 0 {
					recur(i+1, depth+1, sum+row[i]+col[i], flag|(1<<i))
				}
			}
		}
	}
	recur(0, 0, 0, 0)
	fmt.Println(min)
}
