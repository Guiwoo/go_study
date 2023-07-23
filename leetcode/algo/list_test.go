package algo

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Println("node value ", n.val)
	}
}

func TestArray(t *testing.T) {
	var a [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [11]int

	copy(b[0:], a[0:5])
	b[5] = 1000
	copy(b[6:], a[5:])
	fmt.Println(b)
}
