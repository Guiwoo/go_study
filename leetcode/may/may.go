package may

type Point struct {
	row, col int
	dir      []int
}

func generateMatrix(n int) [][]int {
	answer := make([][]int, n)
	for i := range answer {
		answer[i] = make([]int, n)
	}

	cur := &Point{0, 0, []int{0, 1}}

	for i := 0; i < n*n; i++ {
		num := i + 1
		validate(cur, answer)
		fil(cur, answer, num)
	}

	return answer
}

func fil(cur *Point, board [][]int, num int) {
	board[cur.row][cur.col] = num
}

func validate(cur *Point, board [][]int) {
	if board[cur.row][cur.col] == 0 {
		return
	}
	nextRow := cur.row + cur.dir[0]
	nextCol := cur.col + cur.dir[1]

	if nextRow < 0 || nextRow >= len(board) || nextCol < 0 || nextCol >= len(board[0]) || board[nextRow][nextCol] != 0 {
		cur.dir = getNextDir(cur.dir)
	}

	cur.row = cur.row + cur.dir[0]
	cur.col = cur.col + cur.dir[1]
}

func getNextDir(dir []int) (ans []int) {
	switch true {
	case dir[0] == 0 && dir[1] == 1:
		//오른쪽
		ans = []int{1, 0}
	case dir[0] == 1 && dir[1] == 0:
		//아래로
		ans = []int{0, -1}
	case dir[0] == 0 && dir[1] == -1:
		//왼쪽
		ans = []int{-1, 0}
	default:
		ans = []int{0, 1}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	helper(head, true)
	return head
}

func helper(head *ListNode, isEven bool) {
	if head == nil {
		return
	}

	if head.Next != nil && isEven {
		nxt := head
		tmp := head.Next
		nxt.Next = head.Next.Next
		tmp.Next = nxt

		head = tmp
	}

	helper(head.Next, !isEven)
	return
}

func pairSum(head *ListNode) int {
	list := make([]int, 0)

	return getPare(helper2(head, list))
}

func helper2(head *ListNode, list []int) []int {
	if head == nil {
		return list
	}
	return helper2(head.Next, append(list, head.Val))
}

func getPare(list []int) int {
	max := 0
	for i := 0; i < len(list)/2; i++ {
		if max < (list[i] + list[len(list)-i-1]) {
			max = list[i] + list[len(list)-i-1]
		}
	}
	return max
}
