package main

import (
	"fmt"
	"testing"
)

/**
Go 언어 의 동시성 을 매우 쉽게 만ㄷ르어 주기 떄문에 활용하는것이 중요함
Go 의 런타임이 관리하고 생산하는 고루틴이 유익할것으로 판단된다.

작업 가로채기

1. 포크지점 에서 스레드 와 연관 된 데큐의 끝에 작업을 추가한다 . (포크란 ? 분기를 나누는 지점이 되겠다. )
2. 스레드가 유휴 상태이면, 임의의 다른 스레드와 관련된 데큐 앞에서 작업을 가로챈다
3. 아직실현되지 않은 합류지점에서  스레드 자체 데큐 꼬리부분을 제거한다.
4. 스레드 데큐 양쪽이 모두 비어있는경우 .
 - 조인을 지연시키고, 임의의 스레드와 관련된 데큐앞쪽에서 작업을 가로챈다
*/

func Test_runtime_go_01(t *testing.T) {
	var fib func(n int) <-chan int
	fib = func(n int) <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n-1) + <-fib(n-2)
		}()
		return result
	}
	fmt.Printf("fib(4) result : %d ", <-fib(4))
}
