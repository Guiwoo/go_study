package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	a := sc.Text()
	input, _ := strconv.Atoi(a)

	arr := make([]int, input)

	sc.Scan()
	b := strings.Split(sc.Text(), " ")
	for i := range b {
		arr[i], _ = strconv.Atoi(b[i])
	}

	fmt.Println(dfs(arr))
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
