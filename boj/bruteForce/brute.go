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

func p_01() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	list := make([]int, 9)

	var (
		total, a, b int
	)

	for i := range list {
		fmt.Fscanln(reader, &list[i])
		total += list[i]
	}

exit:
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if total-list[i]-list[j] == 100 {
				a = list[i]
				b = list[j]
				break exit
			}
		}
	}

	sort.Ints(list)

	for _, v := range list {
		if v == a || v == b {
			continue
		}
		fmt.Fprintln(writer, v)
	}

	return
}

func p_02() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var list [][]int
	for {
		input, _, _ := reader.ReadLine()
		sub := strings.Split(string(input), " ")

		if sub[0] == "0" {
			break
		}

		subList := make([]int, len(sub)-1)

		for i := 1; i < len(sub); i++ {
			subList[i-1], _ = strconv.Atoi(sub[i])
		}
		sort.Ints(subList)
		list = append(list, subList)
	}

	for _, v := range list {
		generateCombinations(v, 6, []int{}, writer)
		fmt.Fprintln(writer, "")
	}
}

func generateCombinations(arr []int, length int, current []int, writer *bufio.Writer) {
	if len(current) == length {
		for _, num := range current {
			fmt.Fprint(writer, num, " ")
		}
		fmt.Fprintln(writer, "")
		return
	}

	for i, val := range arr {
		if i > len(arr)-(length-len(current)) {
			break
		}
		generateCombinations(arr[i+1:], length, append(current, val), writer)
	}
}

var arr []int
var (
	answer, target int
)

func p_03() {
	sc := bufio.NewScanner(os.Stdin)

	rs := toArr(sc)
	for i := range rs {
		x, _ := strconv.Atoi(rs[i])
		if i == 0 {
			arr = make([]int, x)
		} else {
			target = x
		}
	}

	rs = toArr(sc)
	for i := range rs {
		x, _ := strconv.Atoi(rs[i])
		arr[i] = x
	}
	subArray(0, 0)
	if target == 0 {
		answer--
	}
	fmt.Println(answer)
}

func subArray(idx, sum int) {
	if idx == len(arr) {
		if sum == target {
			answer++
		}
		return
	}
	subArray(idx+1, sum+arr[idx])
	subArray(idx+1, sum)
}

func toArr(sc *bufio.Scanner) []string {
	sc.Scan()
	input := sc.Text()
	rs := strings.Split(input, " ")
	return rs
}

func ps_03_02() {
	sc := bufio.NewScanner(os.Stdin)

	params := toArr(sc)
	arrLen, _ := strconv.Atoi(params[0])
	target, _ := strconv.Atoi(params[1])

	arr := make([]int, arrLen)
	nums := toArr(sc)
	for i, numStr := range nums {
		arr[i], _ = strconv.Atoi(numStr)
	}

	answer := subArray_02(arr, target)
	if target == 0 {
		answer--
	}
	fmt.Println(answer)
}

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
