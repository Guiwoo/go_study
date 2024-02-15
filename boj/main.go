package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

/**
세계적인 도둑 상덕이는 보석점을 털기로 결심했다.
상덕이가 털 보석점에는 보석이 총 N개 있다.
각 보석은 무게 Mi와 가격 Vi를 가지고 있다.
상덕이는 가방을 K개 가지고 있고, 각 가방에 담을 수 있는 최대 무게는 Ci이다.
가방에는 최대 한 개의 보석만 넣을 수 있다.
상덕이가 훔칠 수 있는 보석의 최대 가격을 구하는 프로그램을 작성하시오.

첫째 줄에 N과 K가 주어진다. (1 ≤ N, K ≤ 300,000)
다음 N개 줄에는 각 보석의 정보 Mi와 Vi가 주어진다. (0 ≤ Mi, Vi ≤ 1,000,000)
다음 K개 줄에는 가방에 담을 수 있는 최대 무게 Ci가 주어진다. (1 ≤ Ci ≤ 100,000,000)
모든 숫자는 양의 정수이다.

첫째 줄에 상덕이가 훔칠 수 있는 보석 가격의 합의 최댓값을 출력한다.

2 1
5 10
100 100
11
답 : 10

3 2
1 65
5 23
2 99
10
2

답 :164
*/

type jewel struct {
	Weight int
	Price  int
}

type jewels []*jewel

func (arr jewels) Less(i, j int) bool {
	if arr[i].Price == arr[j].Price {
		return arr[i].Weight < arr[j].Weight
	}
	return arr[i].Price > arr[j].Price
}

func (arr jewels) Len() int {
	return len(arr)
}

func (arr jewels) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr *jewels) Push(x any) {
	item := x.(*jewel)
	*arr = append(*arr, item)
}

func (arr *jewels) Pop() any {
	old := *arr
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*arr = old[0 : n-1]
	return item
}

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		writer = bufio.NewWriter(os.Stdout)
		J, B   int
	)

	defer writer.Flush()
	fmt.Fscanln(reader, &J, &B)

	pq := make(jewels, 0)
	heap.Init(&pq)

	for i := 0; i < J; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		heap.Push(&pq, &jewel{x, y})
	}

	bags := make([]int, 0, B)
	for i := 0; i < B; i++ {
		var x int
		fmt.Fscanln(reader, &x)
		bags = append(bags, x)
	}
	sort.Slice(bags, func(i, j int) bool {
		return bags[i] < bags[j]
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq)
		fmt.Printf("%+v\n", item)
	}

	// 보석을 정렬 가격순으로 무게는 부차적인거
	// 가방은 가벼운 순으로 정렬
	fmt.Printf("%+v\n", bags)

	// 가방 하나 잡고 들어갈수 있을만큼 loop 돌려서 끝내고 넘어가고 answer 추가해주고
	// git test
}
