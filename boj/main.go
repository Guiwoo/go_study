package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
