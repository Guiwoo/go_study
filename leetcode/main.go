package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, num1 := range nums1 {
		for j, num2 := range nums2 {
			if num1 == num2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	answer := maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4})
	fmt.Println(answer)

	answer = maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})
	fmt.Println(answer)

	answer = maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1})
	fmt.Println(answer)

	answer = maxUncrossedLines([]int{1, 2, 3, 4, 5, 6}, []int{6, 2, 3, 4, 5, 1})
	fmt.Println(answer)
}
