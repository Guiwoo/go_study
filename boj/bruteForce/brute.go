package bruteForce

import (
	"bufio"
	"fmt"
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
