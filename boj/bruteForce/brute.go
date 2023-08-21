package bruteForce

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func subArray_02(arr []int, target int) int {
	answer := 0
	n := len(arr)

	// 비트 마스크를 0부터 (2^n - 1)까지 순회합니다.
	for i := 0; i < (1 << n); i++ {
		sum := 0
		// 비트 마스크의 각 비트를 확인하여 포함된 원소의 합을 계산합니다.
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				sum += arr[j]
			}
		}
		if sum == target {
			answer++
		}
	}

	return answer
}

func ps_04() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	input := sc.Text()
	a, _ := strconv.Atoi(input)

	list := make([]int, a)

	sc.Scan()
	input = sc.Text()
	in := strings.Split(input, " ")

	for i := range in {
		list[i], _ = strconv.Atoi(in[i])
	}

	max := dfs(list)
	fmt.Println(max)
}

func dfs(nums []int) int {
	answer := 0
	out := make([]int, len(nums))
	bitmask(nums, out, 0, 0, &answer)
	return answer
}

func bitmask(num, out []int, flag, depth int, answer *int) {
	if depth == len(num) {
		x := getResult(out)
		if x > *answer {
			*answer = x
		}
		return
	}
	for i := 0; i < len(num); i++ {
		if flag&(1<<i) == 0 {
			out[depth] = num[i]
			bitmask(num, out, flag|(1<<i), depth+1, answer)
		}
	}
}

func getResult(num []int) int {
	sum := 0
	for i := 0; i < len(num)-1; i++ {
		sum += int(math.Abs(float64(num[i] - num[i+1])))
	}
	return sum
}

func boj15469(n, m int) string {
	out := make([]int, m)
	sb := strings.Builder{}
	recur(n, 0, 0, out, &sb)
	return sb.String()
}
func recur(n, depth, flag int, out []int, sb *strings.Builder) {
	if depth == len(out) {
		s := strings.NewReplacer("[", "", "]", "").Replace(fmt.Sprintf("%v", out))
		sb.WriteString(s + "\n")
		return
	} else {
		for i := 0; i < n; i++ { // i를 0부터 시작
			if (flag & (1 << i)) != 0 {
				continue
			}
			out[depth] = i + 1 // 실제 숫자는 i + 1
			recur(n, depth+1, flag|(1<<i), out, sb)
		}
	}
}

func boj15650(n, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	recur2(n, 0, 0, 0, out, &sb)
	return sb.String()
}

func recur2(n, start, depth, flag int, out []int, sb *strings.Builder) {
	if depth == len(out) {
		result := strings.NewReplacer("[", "", "]", "").Replace(fmt.Sprintf("%v", out))
		sb.WriteString(result + "\n")
		return
	} else {
		for i := start; i < n; i++ {
			if flag&(1<<i) != 0 {
				continue
			}
			out[depth] = i + 1
			recur2(n, i+1, depth+1, flag|(1<<i), out, sb)
		}
	}
}

/*
*
fmt.Sprintf 또는 NewReplacer 와 같은 문자열 계산 은 상당한 오버헤드가 발생할수 있어 이런 완탐에서 문제가 발생한다.
기억하고 있다가 코드 최적화 시에 적용하자.
*/
func boj15651(n, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	permutation(n, 0, out, &sb)
	return sb.String()
}

func permutation(n, depth int, out []int, sb *strings.Builder) {
	if depth < len(out) {
		for i := 0; i < n; i++ {
			out[depth] = i + 1
			permutation(n, depth+1, out, sb)
		}
	} else {
		for i := 0; i < len(out); i++ {
			sb.WriteString(strconv.Itoa(out[i]))
			if i < len(out)-1 {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
}

/*
*
백준 15652 증가 로 답구하는것
*/
func boj15652(n, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	permutation(n, 0, out, &sb)
	return sb.String()
}
func permutation2(n, depth int, out []int, sb *strings.Builder) {
	if depth == len(out) {
		for _, v := range out {
			sb.WriteString(fmt.Sprintf("%d ", v))
		}
		sb.WriteString("\n")
		return
	} else {
		start := 0
		if depth != 0 {
			start = out[depth-1] - 1
		}
		for i := start; i < n; i++ {
			out[depth] = i + 1
			permutation(n, depth+1, out, sb)
		}
	}
}

/*
*
숫자 주고 중간에 한번 정렬해야함
*/
func boj15654(n []int, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	sort.Ints(n)
	permutation3(n, out, 0, 0, &sb)
	return sb.String()
}

func permutation3(n, out []int, depth, flag int, sb *strings.Builder) {
	if depth == len(out) {
		for i, v := range out {
			if i == len(out)-1 {
				sb.WriteString(fmt.Sprintf("%d", v))
				continue
			}
			sb.WriteString(fmt.Sprintf("%d ", v))
		}
		sb.WriteString("\n")
		return
	} else {
		for i := 0; i < len(n); i++ {
			if flag&(1<<i) != 0 {
				continue
			}
			out[depth] = n[i]
			permutation3(n, out, depth+1, flag|(1<<i), sb)
		}
	}
}

/**
숫자 주고 한번 정렬해서 프린트
*/

func boj15655(n []int, m int) string {
	sb := strings.Builder{}
	out := make([]int, m)
	recur4(n, out, 0, 0, &sb)
	return sb.String()
}

func recur4(n, out []int, depth, flag int, sb *strings.Builder) {
	if depth == len(out) {
		for _, v := range out {
			sb.WriteString(fmt.Sprintf("%d ", v))
		}
		sb.WriteString("\n")
	} else {
		for i := 0; i < len(n); i++ {
			if flag&(1<<i) == 0 {
				if depth == 0 {
					out[depth] = n[i]
					recur4(n, out, depth+1, flag|(1<<i), sb)
				} else if out[depth-1] < n[i] {
					out[depth] = n[i]
					recur4(n, out, depth+1, flag|(1<<i), sb)
				}
			}
		}
	}
}

var buf bytes.Buffer

func boj15656(n []int, m int) string {
	out := make([]int, m)
	boj15656Recursion(n, out, 0)
	return buf.String()
}

func boj15656Recursion(n, out []int, depth int) {
	if depth == len(out) {
		for i := range out {
			buf.WriteString(fmt.Sprintf("%d ", out[i]))
		}
		buf.WriteByte('\n')
		return
	} else {
		for i := 0; i < len(n); i++ {
			out[depth] = n[i]
			boj15656Recursion(n, out, depth+1)
		}
	}
}

func boj15657(n []int, m int) {
	out := make([]int, m)
	boj15657Recur(n, out, 0)
	fmt.Println(buf.String())
}
func boj15657Recur(n, out []int, depth int) {
	if depth == len(out) {
		for _, v := range out {
			buf.WriteString(strconv.Itoa(v) + " ")
		}
		buf.WriteByte('\n')
		return
	}

	for i := 0; i < len(n); i++ {
		if depth == 0 || out[depth-1] <= n[i] {
			out[depth] = n[i]
			boj15657Recur(n, out, depth+1)
		}
	}
}

var (
	out     []int
	checker map[string]bool
	sb      strings.Builder
)

func boj15663(arr []int, m int) {
	out = make([]int, m)
	checker = make(map[string]bool)
	sb = strings.Builder{}
	recurBoj15663(arr, 0, 0)
	fmt.Printf(sb.String())
}

// 1nt64 왜 int 값을 벗어날까 ? *10 을 해서 ? ㅂㅅ 임
func recurBoj15663(arr []int, depth, flag int) {
	if depth == len(out) {
		var (
			target       int
			targetString string
		)
		for _, v := range out {
			target = target*10 + v
			targetString += fmt.Sprintf("%d ", v)
		}
		if !checker[targetString] {
			checker[targetString] = true
			sb.WriteString(fmt.Sprintf("%s\n", targetString))
		}
		return
	} else {
		for i := 0; i < len(arr); i++ {
			if flag&(1<<i) == 0 {
				out[depth] = arr[i]
				recurBoj15663(arr, depth+1, flag|(1<<i))
			}
		}
	}
}

func boj15664(arr []int, m int) {
	out := make([]int, m)
	sb := strings.Builder{}
	checker := make(map[string]bool)

	recur15664(arr, out, 0, 0, checker, &sb)
	fmt.Println(sb.String())
}

func recur15664(arr, out []int, depth, flag int, checker map[string]bool, sb *strings.Builder) {
	if depth == len(out) {
		var (
			targetSt string
			prev     int
		)
		for _, v := range out {
			targetSt += fmt.Sprintf("%d ", v)
			if prev > v {
				return
			}
			prev = v
		}

		if !checker[targetSt] {
			checker[targetSt] = true
			sb.WriteString(fmt.Sprintf("%s\n", targetSt))
		}
		return
	} else {
		for i := 0; i < len(arr); i++ {
			if flag&(1<<i) == 0 {
				out[depth] = arr[i]
				recur15664(arr, out, depth+1, flag|(1<<i), checker, sb)
			}
		}
	}
}

func boj15665(arr []int, m int) {
	out := make([]int, m)
	sb := strings.Builder{}
	checker := make(map[string]bool)

	recur15665(arr, out, 0, checker, &sb)
	fmt.Println(sb.String())
}

func recur15665(arr, out []int, depth int, checker map[string]bool, sb *strings.Builder) {
	if depth == len(out) {
		var (
			targetSt string
		)
		for _, v := range out {
			targetSt += fmt.Sprintf("%d ", v)
		}

		sb.WriteString(fmt.Sprintf("%s\n", targetSt))
		return
	} else {
		var pre int
		for i := 0; i < len(arr); i++ {
			if pre != arr[i] {
				out[depth] = arr[i]
				pre = arr[i]
				recur15665(arr, out, depth+1, checker, sb)
			}
		}
	}
}

var (
	mapper map[string]bool
)

func boj15666(arr []int, m int) string {
	out := make([]int, m)
	sb := strings.Builder{}
	mapper = make(map[string]bool)
	recur15666(arr, out, 0, &sb)
	return sb.String()
}

func recur15666(arr, out []int, depth int, sb *strings.Builder) {
	if depth == len(out) {
		s := ""
		for _, v := range out {
			s += fmt.Sprintf("%d ", v)
		}
		if !mapper[s] {
			mapper[s] = true
			sb.WriteString(s + "\n")
		}
		return
	} else {
		for i := 0; i < len(arr); i++ {
			if depth == 0 {
				out[depth] = arr[i]
				recur15666(arr, out, depth+1, sb)
			}
			if depth > 0 && out[depth-1] <= arr[i] {
				out[depth] = arr[i]
				recur15666(arr, out, depth+1, sb)
			}
		}
	}
}

func boj15661(arr [][]int) {
	var (
		ans   = math.MaxInt
		n     = len(arr)
		recur func(idx, bit, cnt int)
	)
	min := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	recur = func(idx, bit, cnt int) {
		if cnt > n/2 {
			return
		}
		var (
			a, b int
		)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if bit&(1<<i) == 0 && bit&(1<<j) == 0 {
					a += arr[i][j]
					a += arr[j][i]
				} else if bit&(1<<i) != 0 && bit&(1<<j) != 0 {
					b += arr[i][j]
					b += arr[j][i]
				}
			}
		}
		if ans > min(a, b) {
			ans = min(a, b)
		}
		for i := idx + 1; i < n; i++ {
			recur(i, bit|(1<<i), cnt+1)
		}
	}

	for i := 0; i < len(arr); i++ {
		recur(i, 1<<i, 0)
	}
	fmt.Println(ans)
}

func boj15661Ver2(row, col []int, total int) {
	var answer = total
	var recur func(a, b int)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	recur = func(depth int, sum int) {
		answer = min(answer, abs(sum))
		for i := depth + 1; i < len(row); i++ {
			recur(i, sum-row[i]-col[i])
		}
	}

	answer = min(total, answer)
	recur(0, total)

	fmt.Println(answer)
}

func boj2580(arr [][]int) {
	dfs01(arr, 0)
}

func check(arr [][]int, row, col, i int) bool {
	// 가로줄 검사
	for j := 0; j < 9; j++ {
		if arr[row][j] == i {
			return false
		}
	}

	// 세로줄 검사
	for j := 0; j < 9; j++ {
		if arr[j][col] == i {
			return false
		}
	}

	// 3*3 검사
	row = (row / 3) * 3
	col = (col / 3) * 3
	for j := row; j < row+3; j++ {
		for k := col; k < col+3; k++ {
			if arr[j][k] == i {
				return false
			}
		}
	}
	return true
}

func dfs01(arr [][]int, cur int) {
	row := cur / 9
	col := cur % 9

	if cur == 81 {
		for i := range arr {
			for j := range arr[i] {
				fmt.Printf("%d ", arr[i][j])
			}
			fmt.Println()
		}
		os.Exit(0)
	}

	if arr[row][col] == 0 {
		for i := 1; i <= 9; i++ {
			if check(arr, row, col, i) {
				arr[row][col] = i
				dfs01(arr, cur+1)
				arr[row][col] = 0
			}
		}
	} else {
		dfs01(arr, cur+1)
	}
}

var minResult, maxResult string

func boj2529(arr []string) {
	maxResult, minResult = "", "999999999"
	for i := 0; i <= 9; i++ {
		boj2529Recur(arr, []int{i}, 0, 1<<i)
	}
	fmt.Println(maxResult)
	fmt.Println(minResult)
}

func boj2529Recur(arr []string, result []int, idx, flag int) {
	if len(result) == len(arr)+1 {
		var str string
		for _, v := range result {
			str += fmt.Sprintf("%d", v)
		}
		if str > maxResult {
			maxResult = str
		}
		if str < minResult {
			minResult = str
		}
		return
	} else {
		if arr[idx] == "<" {
			for i := result[idx] + 1; i <= 9; i++ {
				if flag&(1<<i) == 0 {
					boj2529Recur(arr, append(result, i), idx+1, flag|(1<<i))
				}
			}
		} else {
			for i := result[idx] - 1; i >= 0; i-- {
				if flag&(1<<i) == 0 {
					boj2529Recur(arr, append(result, i), idx+1, flag|(1<<i))
				}
			}
		}
	}
}

func boj14888(arr, oper []int) (int, int) {
	max, min := -1000000000, 1000000000
	boj14888Recur(arr, oper, 0, arr[0], &max, &min)
	return max, min
}

func boj14888Recur(arr, oper []int, idx, rs int, max, min *int) {
	if idx == len(arr)-1 {
		if rs > *max {
			*max = rs
		}
		if rs < *min {
			*min = rs
		}
		return
	}
	for i := 0; i < len(oper); i++ {
		if oper[i] > 0 {
			oper[i]--
			nextVal := rs
			switch i {
			case 0:
				nextVal += arr[idx+1]
			case 1:
				nextVal -= arr[idx+1]
			case 2:
				nextVal *= arr[idx+1]
			default:
				nextVal /= arr[idx+1]
			}
			boj14888Recur(arr, oper, idx+1, nextVal, max, min)
			oper[i]++
		}
	}
}

func boj15686() {
	answer := 1 << 31
	var distance [][]int
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minValue := func(include []int, distance [][]int) int {
		result := 0
		for i := 0; i < len(distance); i++ {
			step := 1 << 31
			for j := 0; j < len(distance[i]); j++ {
				for k := 0; k < len(include); k++ {
					if j == include[k] {
						step = min(step, distance[i][j])
					}
				}
			}
			result += step
		}
		return result
	}

	helper := func(num, depth, limit, flag, start int, out []int) {}
	helper = func(num, depth, limit, flag, start int, out []int) {
		if depth == limit {
			answer = min(answer, minValue(out, distance))
			return
		} else {
			for i := start; i < num; i++ {
				if flag&(1<<i) == 0 {
					helper(num, depth+1, limit, flag|(1<<i), i+1, append(out, i))
				}
			}
		}
	}
	reader := bufio.NewReader(os.Stdin)

	var (
		a, b int
	)
	fmt.Fscanln(reader, &a, &b)
	house := make([][]int, 0)
	chicken := make([][]int, 0)

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			var tmp int
			fmt.Fscan(reader, &tmp)
			if tmp == 1 {
				house = append(house, []int{i, j})
			} else if tmp == 2 {
				chicken = append(chicken, []int{i, j})
			}
		}
	}

	distance = make([][]int, len(house))
	for i := 0; i < len(house); i++ {
		distance[i] = make([]int, len(chicken))
		for j := 0; j < len(chicken); j++ {
			distance[i][j] = abs(house[i][0]-chicken[j][0]) + abs(house[i][1]-chicken[j][1])
		}
	}

	/**
	고민해야할 사항 : 치킨집의 개수가 최대 13개이므로, 치킨집의 개수에서 M개를 뽑는 조합을 구하면 된다. (13C6) = 1716 * 집이 최대 50개이므로, 1716 * 50 = 85800
	*/
	helper(len(chicken), 0, b, 0, 0, []int{})

	fmt.Println(answer)
}

func boj14500() {
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
