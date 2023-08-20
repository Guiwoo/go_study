package data_structure

import (
	"fmt"
	"testing"
)

/**
"한정된 범위" 를 가지는 정수들의 배열을 정렬하는 알고리즘
1. 한정된 범위 라는 전제조건 때문에 사용할 수 있는 알고리즘
2. 제한조건 맞는경우 제일 빠른 정렬 알고리즘
Quick sort => 0(N), O(NlogN), O(N^2)
Merge sort => O(NlogN) 일정한 속도 유지
Counting sort => O(N+K)
*/

func Test_CountingSort(t *testing.T) {
	type students struct {
		Name   string
		Height float64
	}
	countSort := func(arr []int, max int) []int {
		result := make([]int, 0, len(arr))

		marker := make([]int, max+1)
		for _, v := range arr {
			marker[v] += 1
		}

		for i, v := range marker {
			for j := 0; j < v; j++ {
				result = append(result, i)
			}
		}
		return result
	}

	countSortBetter := func(arr []int, max int) []int {
		result := make([]int, len(arr))

		marker := make([]int, max+1)
		for _, v := range arr {
			marker[v]++
		}
		for i := range marker {
			if i != 0 {
				marker[i] += marker[i-1]
			}
		}
		for i := range arr {
			result[marker[arr[i]]-1] = arr[i]
			marker[arr[i]]--
		}
		return result
	}

	//알파벳 소문자로 이루어져있는 알파벳중 가장 많이 나오는 알파벳 케릭터 출력
	// 데이터가 스트림 형태로 들어온다면 어떻게 처리할래 ? => 힙데이터 구조를 이용해야함
	alpha := func(input string) string {
		alphabet := make([]int, 26)
		for _, v := range input {
			alphabet['a'-v]++
		}
		var idx = 0
		for i, v := range alphabet {
			if alphabet[idx] < v {
				idx = i
			}
		}
		return string(byte('a' + idx))
	}

	filterHeight := func(a, b int, arr []students) {
		var heightMap [3000][]string
		for _, v := range arr {
			idx := int(v.Height * 10)
			heightMap[idx] = append(heightMap[idx], v.Name)
		}
		for i := a * 10; i < b*10; i++ {
			for _, name := range heightMap[i] {
				fmt.Printf("name %s , height %.1f", name, float32(i)/10)
			}
		}
	}

	arr := []int{4, 2, 3, 1, 2, 2, 1, 1, 1, 1, 1, 6, 4, 3, 7, 9, 8, 1, 2, 3, 7, 6, 9}

	fmt.Println("count sort : ", countSort(arr, 9))
	fmt.Println("count sort : ", countSortBetter(arr, 9))
	fmt.Println("alphabet find : ", alpha("aaaaaaaaaa"))
	stu := []students{
		{"jay", 223.9},
		{"m", 123.4},
		{"ab", 167.7},
		{"abcd", 190.7},
		{"joejoe", 177.3},
		{"holymoly", 175.2},
	}
	fmt.Println("student")
	filterHeight(160, 170, stu)
}
