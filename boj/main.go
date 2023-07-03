package main

import "fmt"

func main() {
	x := solution2([]int{1, 2, 3, 4, 5}, 7)
	fmt.Println(x)
}

func solution(sequence []int, k int) []int {

	var answer []int

	for i := 0; i < len(sequence); i++ {
		if sequence[i] == k {
			return []int{i, i}
		}
		left := i
		sum := sequence[i]
		for j := i + 1; j < len(sequence); j++ {
			sum += sequence[j]
			if sum == k {
				if len(answer) > 0 {
					// 비교 해서 넣어주기
					cur := answer[1] - answer[0]
					could := j - i
					if could < cur {
						answer = []int{i, j}
					}
				} else {
					answer = []int{i, j}
				}
			} else if sum > k {
				sum = -sequence[left]
				left++
			}
		}
	}
	return answer
}

func solution2(seq []int, k int) []int {
	prep := make([]int, len(seq))
	prep[0] = seq[0]
	for i := 1; i < len(seq); i++ {
		prep[i] = seq[i] + prep[i-1]
	}
	var answer []int
	for i := 0; i < len(prep); i++ {
		if prep[i] == k {
			return []int{0, i}
		}
		if prep[i] > k {
			//앞에 빼줘야하잖아
			for j := 0; j < i; j++ {
				target := prep[i] - prep[j]
				if target == k {
					if len(answer) > 0 {
						cur := answer[1] - answer[0]
						now := i - j
						if now < cur {
							answer = []int{j + 1, i}
						}
					} else {
						answer = []int{j + 1, i}
						break
					}
				} else if target < k {
					break
				}
			}
		}
	}
	return answer
}

func sol3(seq []int, k int) []int {
	idx := 0
	sum := seq[0]
	var answer []int

	for i := 1; i < len(seq); i++ {
		sum += seq[i]
		if sum == k {
			answer = getAnswer(answer, i, idx)
		} else if sum > k {
			for sum > k && idx < i {
				sum -= seq[idx]
				idx++
			}

			if sum == k {
				answer = getAnswer(answer, i, idx)
			}
		}
	}
	return answer
}

func getAnswer(answer []int, i int, idx int) []int {
	if len(answer) > 0 {
		cur := answer[1] - answer[0]
		now := i - idx
		if now < cur {
			answer = []int{idx, i}
		}
	} else {
		answer = []int{idx, i}
	}
	return answer
}
