package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(input, " ")
	n, _ := strconv.Atoi(inputs[0])
	m, _ := strconv.Atoi(strings.TrimSpace(inputs[1]))

	result := boj15651(n, m)
	writer.WriteString(result)
}
