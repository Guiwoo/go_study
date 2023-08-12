package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

/**
5 3
0 0 1 0 0
0 0 2 0 1
0 1 2 0 0
0 0 1 0 0
0 0 0 0 2
*/

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	arr := make([][]int, n)

	fmt.Println(n, m)
}
