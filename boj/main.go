package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/**
한 개의 회의실이 있는데 이를 사용하고자 하는 N개의 회의에 대하여 회의실 사용표를 만들려고 한다.
각 회의 I에 대해 시작시간과 끝나는 시간이 주어져 있고, 각 회의가 겹치지 않게 하면서 회의실을 사용할 수 있는 회의의 최대 개수를 찾아보자.
단, 회의는 한번 시작하면 중간에 중단될 수 없으며 한 회의가 끝나는 것과 동시에 다음 회의가 시작될 수 있다.
회의의 시작시간과 끝나는 시간이 같을 수도 있다.
이 경우에는 시작하자마자 끝나는 것으로 생각하면 된다.

첫째 줄에 회의의 수 N(1 ≤ N ≤ 100,000)이 주어진다.
둘째 줄부터 N+1 줄까지 각 회의의 정보가 주어지는데 이것은 공백을 사이에 두고 회의의 시작시간과 끝나는 시간이 주어진다.
시작 시간과 끝나는 시간은 231-1보다 작거나 같은 자연수 또는 0이다.

첫째 줄에 최대 사용할 수 있는 회의의 최대 개수를 출력한다.

11
1 4
3 5
5 7
12 14
0 6
2 13
3 8
5 9
6 10
8 11
8 12

4
*/

type Meeting struct {
	start, end int
}

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		N      int
	)

	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	arr := make([]Meeting, 0, N)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		arr = append(arr, Meeting{x, y})
	}

	sort.Slice(arr, func(i, j int) bool {
		dis1, dis2 := arr[i].end-arr[i].start, arr[j].end-arr[j].start
		if arr[i].end < arr[j].end {
			return dis1 < dis2
		}
		return arr[i].end < arr[j].end
	})

	fmt.Printf("%+v", arr)
}
