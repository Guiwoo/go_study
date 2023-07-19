package july

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return res
}

func romanToInt(s string) int {
	m := make(map[int]int)
	setValue := func(m map[int]int) {
		m['I'] = 1
		m['V'] = 5
		m['X'] = 10
		m['L'] = 50
		m['C'] = 100
		m['D'] = 500
		m['M'] = 1000
	}
	setValue(m)

	var answer int
	var i int

	for i < len(s)-1 {
		idx := int(s[i])
		next := int(s[i+1])
		if m[idx] < m[next] {
			answer += m[next] - m[idx]
			i++
		} else {
			answer += m[idx]
		}
		i++
		fmt.Println(answer)
	}

	if i != len(s) {
		answer += m[int(s[len(s)-1])]
	}
	return answer
}
