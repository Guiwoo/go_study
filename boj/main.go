package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
n(2 ≤ n ≤ 100)개의 도시가 있다. 그리고 한 도시에서 출발하여 다른 도시에 도착하는 m(1 ≤ m ≤ 100,000)개의 버스가 있다.
각 버스는 한 번 사용할 때 필요한 비용이 있다.
모든 도시의 쌍 (A, B)에 대해서 도시 A에서 B로 가는데 필요한 비용의 최솟값을 구하는 프로그램을 작성하시오.

첫째 줄에 도시의 개수 n이 주어지고 둘째 줄에는 버스의 개수 m이 주어진다.
그리고 셋째 줄부터 m+2줄까지 다음과 같은 버스의 정보가 주어진다.
먼저 처음에는 그 버스의 출발 도시의 번호가 주어진다.
버스의 정보는 버스의 시작 도시 a, 도착 도시 b, 한 번 타는데 필요한 비용 c로 이루어져 있다.
시작 도시와 도착 도시가 같은 경우는 없다. 비용은 100,000보다 작거나 같은 자연수이다.

시작 도시와 도착 도시를 연결하는 노선은 하나가 아닐 수 있다.

5
14
1 2 2
1 3 3
1 4 1
1 5 10
2 4 2
3 4 1
3 5 1
4 5 3
3 5 10
3 1 8
1 4 2
5 1 7
3 4 2
5 2 4
*/

func main() {
	boj11404()
}

func boj11404() {
	var (
		reader      = bufio.NewReader(os.Stdin)
		writer      = bufio.NewWriter(os.Stdout)
		cities, bus int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &cities)
	fmt.Fscanln(reader, &bus)

	graph := make([]map[int][]int, cities+1)
	for i := range graph {
		graph[i] = make(map[int][]int)
	}

	for i := 0; i < bus; i++ {
		var from, to, value int
		fmt.Fscanln(reader, &from, &to, &value)
		graph[from][to] = append(graph[from][to], value)
	}

	for i := 1; i <= cities; i++ {
		way := findWay(i, cities, graph)
		fmt.Fprintln(writer, way)
	}
}

func findWay(start, cities int, graph []map[int][]int) string {
	dp := make([]int, cities+1)
	for i := range dp {
		dp[i] = 1e8
	}
	dp[start] = 0

	for i := 1; i <= cities; i++ {
		for to, values := range graph[i] {
			for _, value := range values {
				if dp[to] > dp[i]+value {
					dp[to] = dp[i] + value
				}
			}
		}
	}

	for to, values := range graph[1] {
		for _, value := range values {
			if dp[to] > dp[1]+value {
				dp[to] = dp[1] + value
			}
		}
	}

	answer := ""
	for i := 1; i < len(dp); i++ {
		if dp[i] == 1e8 {
			dp[i] = 0
		}
		answer += fmt.Sprintf("%d ", dp[i])
	}

	return answer
}
