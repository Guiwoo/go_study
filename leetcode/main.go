package main

import (
	"sort"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
어떻게 루프를 태워야 할까 누구를 기준으로 돌려야할지 모르겠음 각 스트링 별로 루프 태워 ? prefix 잖아 앞에서부터 맞춰야지
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]
	var prefix strings.Builder

	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			break
		}
		prefix.WriteByte(first[i])
	}

	return prefix.String()
}

func main() {
	str := []string{"cir", "car"}
	longestCommonPrefix(str)
}
