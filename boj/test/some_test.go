package test

import (
	"container/heap"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_stringBuilder(t *testing.T) {
	sb := strings.Builder{}

	for i := 0; i < 3; i++ {
		sb.WriteString("something")
	}
	fmt.Println(sb.String())
}

func Test_bitMastk(t *testing.T) {
	fmt.Println(1 << 0)
	fmt.Println(0 | (1 << 0))
	fmt.Println(0 & (1 << 0))
}

func Test_StringSplit(t *testing.T) {
	a := "abcdefg"
	fmt.Println(a[:1])
	fmt.Println(a[1:3])
	fmt.Println(a[3:])
	fmt.Println(a[7:])
}

func Test_Recur(t *testing.T) {
	recur := func(n int, rs []int, start, depth, flag int) {}
	recur = func(n int, rs []int, start, depth, flag int) {
		if depth == 2 {
			fmt.Println(rs)
			return
		} else {
			for i := start; i < n; i++ {
				if (flag & (1 << i)) == 0 {
					recur(n, append(rs, i+1), i+1, depth+1, flag|(1<<i))
				}
			}
		}
	}

	recur2 := func(n int, rs []int, flag, depth int) {}
	recur2 = func(n int, rs []int, flag, depth int) {
		if depth >= 2 && len(rs) == 2 {
			fmt.Println(rs)
			return
		} else {
			for i := 0; i < n; i++ {
				if flag&(1<<i) == 0 {
					recur2(n, append(rs, i+1), flag|(1<<i), depth+1)
				}
			}
		}
	}

	recur(4, []int{}, 0, 0, 0)
	fmt.Println("======================")
	recur2(4, []int{}, 0, 0)
}

func Test_Recur_loop(t *testing.T) {
	helper := func(a, depth, start, flag, result int, out string) {}
	helper = func(a, depth, start, flag, result int, out string) {
		if depth == result {
			fmt.Println(out)
			return
		} else {
			for i := start; i < a; i++ {
				if flag&(1<<i) == 0 {
					helper(a, depth+1, i+1, flag|(1<<i), result, out+fmt.Sprintf("%d ", i))
				}
			}
		}
	}

	for i := 1; i <= 2; i++ {
		helper(4, 0, 0, 0, i, "")
		fmt.Println("===============")
	}
}

func TestBoj(t *testing.T) {
	aFunc := func(arr []int) []int {
		fmt.Printf("%p\n", arr)
		arr[0] = 9999
		return arr
	}

	arr := []int{123}

	arr2 := aFunc(arr)
	fmt.Println(arr, arr2)
	fmt.Printf("%p,%p", arr, arr2)
}

func TestEqual(t *testing.T) {
	type tester struct {
		a int
		b string
		c bool
	}
	a := tester{5, "a", true}
	b := tester{5, "a", true}
	fmt.Println("ab")
	if a == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println("cd")

	c := []int{1, 2, 3}
	d := []int{1, 2, 3}
	if reflect.DeepEqual(c, d) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func solution(ball, order []int) []int {
	left, right := 0, len(ball)-1
	arr := make(map[int]bool)
	answer := make([]int, 0, len(ball))

	process := func() {
		for {
			if left <= right && arr[ball[left]] {
				delete(arr, ball[left])
				answer = append(answer, ball[left])
				left++
			} else if left <= right && arr[ball[right]] {
				delete(arr, ball[right])
				answer = append(answer, ball[right])
				right--
			} else {
				break
			}
		}
	}

	for _, o := range order {
		process()
		if o == ball[left] {
			answer = append(answer, ball[left])
			left++
		} else if o == ball[right] {
			answer = append(answer, ball[right])
			right--
		} else {
			arr[o] = true
		}
	}
	process() // 마지막으로 남은 요소 처리

	return answer
}

type Room struct {
	Number      int
	max, people int
}

type PriorityQueue []*Room

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// 방이 max 이상 찬 경우 가장 후순위로 설정
	if pq[i].people >= pq[i].max && pq[j].people >= pq[j].max {
		return pq[i].Number < pq[j].Number
	}
	if pq[i].people >= pq[i].max {
		return false
	}
	if pq[j].people >= pq[j].max {
		return true
	}
	// people이 같으면 Number 순서대로 정렬
	if pq[i].people == pq[j].people {
		return pq[i].Number < pq[j].Number
	}
	return pq[i].people > pq[j].people
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Room)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(roomNum, change int) {
	idx := -1
	for i := range *pq {
		if (*pq)[i].Number == roomNum {
			idx = i
			break
		}
	}

	if idx == -1 {
		return
	}

	(*pq)[idx].people += change

	heap.Fix(pq, idx)
}

func Test03(t *testing.T) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Room{Number: 1, max: 2, people: 1})
	heap.Push(&pq, &Room{Number: 2, max: 2, people: 2})
	heap.Push(&pq, &Room{Number: 3, max: 2, people: 1})
	heap.Push(&pq, &Room{Number: 4, max: 2, people: 1})
	pq.update(4, -1)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Room)
		fmt.Printf("%+v\n", item)
	}
}

func solution2(n int, entry []int) []int {
	//cur := 1
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Room{1, n, 0})

	for i := 0; i < len(entry); i++ {
		// heap 의 첫값이 max 랑 people 랑 같으면 ? 방생성해서 넣어주기
		//c := heap.Pop(&pq).(*Room)
		//if entry[i] == 0 {
		//	c.people++
		//} else {
		//	c.people--
		//}
		//heap.Push(&pq, c)
	}
	return nil
}

func Test01(t *testing.T) {
	type input struct {
		name   string
		ball   []int
		order  []int
		answer []int
	}
	inputs := []input{
		{
			"예제 1",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{6, 2, 5, 1, 4, 3},
			[]int{6, 5, 1, 2, 4, 3},
		},
		{
			"예제 2",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{6, 2, 5, 1, 4, 3},
			[]int{6, 5, 1, 2, 4, 3},
		},
		{
			"예제 3",
			[]int{11, 2, 9, 13, 24},
			[]int{9, 2, 13, 24, 11},
			[]int{24, 13, 9, 2, 11},
		},
		{
			"1자리 수",
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			"순차적입력",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, i := range inputs {
		t.Run(i.name, func(t *testing.T) {
			answer := solution(i.ball, i.order)
			for idx := range answer {
				if answer[idx] != i.answer[idx] {
					t.Errorf("fail testing got %+v expected %+v", answer, i.answer)
					break
				}
			}
		})
	}
}
