package shortest_path

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	TC      int
	N, M, W int
	adjs    [][]Adj
	upper   []int
)

type Adj struct {
	node   int
	weight int
}

const INF = 987654321

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	TC = scanInt()
	for i := 1; i <= TC; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N, M, W = scanInt(), scanInt(), scanInt()

	adjs = make([][]Adj, N+1)

	for i := 1; i <= M; i++ {
		s, e, t := scanInt(), scanInt(), scanInt()
		adjs[s] = append(adjs[s], Adj{e, t})
		adjs[e] = append(adjs[e], Adj{s, t})
	}

	for i := 1; i <= W; i++ {
		s, e, t := scanInt(), scanInt(), scanInt()
		adjs[s] = append(adjs[s], Adj{e, -t})
	}

	upper = make([]int, N+1)
	for i := 1; i <= N; i++ {
		upper[i] = INF
	}
	upper[1] = 0
}

func Solve() {
	if hasNegativeCycle() {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func hasNegativeCycle() bool {
	for i := 1; i <= N-1; i++ {
		if !relax() {
			return false
		}
	}
	return relax()
}

func relax() bool {
	relaxed := false
	for from := 1; from <= N; from++ {
		for _, adj := range adjs[from] {
			to := adj.node
			weight := adj.weight

			if upper[to] > upper[from]+weight {
				upper[to] = upper[from] + weight
				relaxed = true
			}
		}
	}
	return relaxed
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
