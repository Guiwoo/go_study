package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/**
20
7
23
19
10
15
25
8
13
*/

func main() {
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
		fmt.Fprintln(writer, list[i])
	}

	return
}
