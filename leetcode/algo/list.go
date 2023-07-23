package algo

type Node[T any] struct {
	next *Node[T]
	val  T
}

func (n *Node[T]) Append(next *Node[T]) *Node[T] {
	n.next = next
	return n
}

/**
Linked List
불연속 메모리 - 필요한 만큼 메모리 사용
삽입/삭제 에 효율적
Random access o(n) 이다
cache miss 가 잘일어남
요소가 사라질때 GC 발생
*/

/**
Array
연속된 메모리 한번에 생성, 메모리 해제
Random Access 에 매우효율적 => 변수의 시작메모리 + 인덱스 * 타입사이즈
삽입/삭제에 약하다. 배열 끝 추가와 삭제는 okay
cache miss 가 잘 일어나지 않는다(cpu 의 캐시가 가져오는것) 고도의 집약된 데이터 => array 가 유리
요소가 사라질떄 마다 GC 되지 않는다 !!
*/
