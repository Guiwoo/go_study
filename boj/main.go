package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution()
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var n int
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		fmt.Fprintln(writer, boj3933(n))
	}
}

func boj3933(n int) int {
	var cnt int
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			cnt++
			continue
		} else if i*i > n {
			break
		}
		for j := i; i*i+j*j <= n; j++ {
			if i*i+j*j == n {
				cnt++
				continue
			} else if i*i+j*j > n {
				break
			}
			for k := j; i*i+j*j+k*k <= n; k++ {
				if i*i+j*j+k*k == n {
					cnt++
					continue
				} else if i*i+j*j+k*k > n {
					break
				}
				for l := k; i*i+j*j+k*k+l*l <= n; l++ {
					if i*i+j*j+k*k+l*l == n {
						cnt++
						continue
					} else if i*i+j*j+k*k+l*l > n {
						break
					}
				}
			}
		}

	}
	return cnt
}
