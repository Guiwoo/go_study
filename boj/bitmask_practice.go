package main

import "fmt"

/*
*
비트 마스킹을 이용한 홀짝 판별법
짝수라면 true, 홀수라면 false
*/
func oddEvenChecker(a int) bool {
	fmt.Println(a & 1) //가장 오른쪽 비트 기준으로 연산되기 때문
	return 0 == a&1
}

/*
*
특정 비트 토글 주어진 정수 n 의 k번쨰 비트를 반전하라
*/
func toggleBitCheck(n, k int) int {
	return n ^ (1 << k)
}

/*
*
방문 체크 비트 마스킹
*/
func checkFlag(arr []bool, flag int) {
	for i := range arr {
		if flag&(1<<i) != 0 {
			continue
		}
		arr[i] = true
		checkFlag(arr, flag|(1<<i))
	}
}
