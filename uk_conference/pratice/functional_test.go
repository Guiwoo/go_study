package praticea

import (
	"fmt"
	"testing"
)

func mapFunc[T any](slice []T, fn func(T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func TestMultiply(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := mapFunc(input, func(x int) int {
		return x * 2
	})
	fmt.Println(result) // 예상 출력: [2, 4, 6, 8]

}
