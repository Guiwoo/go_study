package main

import (
	"bufio"
	"os"
)

func main() {
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	writer.WriteString("")
}
