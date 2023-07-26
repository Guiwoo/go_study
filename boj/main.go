package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution()
}

/*
*
fsacn 을 활용한 읽기 방법 다른사람의 방식
reader := bufio.NewReader(os.Stdin)

defer writer.Flush()

var N, M int
fmt.Fscanln(reader, &N, &M)

values = make([]int, N)

	for i := range values{
		fmt.Fscan(reader, &values[i])
	}

sort.Ints(values)
*/
func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n []int
		m int
	)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(input, " ")
	length, _ := strconv.Atoi(inputs[0])
	m, _ = strconv.Atoi(strings.TrimSpace(inputs[1]))

	n = make([]int, length)

	input, _ = reader.ReadString('\n')
	inputs = strings.Split(strings.TrimSpace(input), " ")
	for i, v := range inputs {
		n[i], _ = strconv.Atoi(v)
	}
	//writer.WriteString(boj15654(n, m))
	writer.WriteString(fmt.Sprintf("%d", m))
}
