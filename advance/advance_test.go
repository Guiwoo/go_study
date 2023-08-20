package main

import (
	"fmt"
	"testing"
)

func Test_String(t *testing.T) {
	a := []string{"apple", "banana", "citron"}

	for _, v := range a {
		v = "bullshit"
		fmt.Println(v)
	}
	fmt.Println(a)
}

func Test_Array_Address(t *testing.T) {
	arr := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("Array Address : %p\n", &arr)
	for i, v := range arr {
		fmt.Printf("Value : %s,Value Address: %p, Address : %p\n", v, &v, &arr[i])
	}
	/**
	fmt.Printf("Value : %s,Value Address: %p, Address : %p\n", v, &v, &arr[i])
	Value : a,Value Address: 0xc0000a0000, Address : 0xc0000a0000
	Value : b,Value Address: 0xc0000a0000, Address : 0xc0000a0010
	Value : c,Value Address: 0xc0000a0000, Address : 0xc0000a0020

	v 의 주소는 동일한 값을 사용하는데 왜 ? 배열은 동일한 타입이기 때문에 길이가 같아 해당 주소를 계속 사용하며 스택에 쌓아올릴수 있다.
	&arr[i] 의 주소는 10 씩 증가해서 음 10바이트 씩 증가하네 ? 스트링은 16바이트 (포인터 8바이트, 길이 8바이트)여야 하는데 ? 라고 생각이 들수 있다.
	16진수에서 10 이 16이다.
	*/
}

func Test_Slice_Address(t *testing.T) {
	arr := make([]string, 6)
	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	arr[3] = "d"
	arr[4] = "e"
	arr[5] = "f"

	fmt.Printf("Slice name Address : %p\n", &arr)
	for i := range arr {
		fmt.Printf("Slice Value Address : %p\n", &arr[i])
	}

	/**
	arr 라는 변수가 가지는 주소와 arr[i] 의 0번 주소는 다르다.
	왜 ? 다를까 ? 슬라이스는 구조체 이다. 따라서 슬라이스의 주소값을 찍으면 그건 구조체의 주소값이다.
	*/
}

type Guiwoo struct {
	name   string
	age    int
	height int
	job    string
}

func Test_Slice_nil_empty(t *testing.T) {
	var data []string
	data2 := []string{}
	if data == nil {
		fmt.Println("data is nil")
	} else {
		fmt.Println("data is not nil")
	}

	if data2 == nil {
		fmt.Println("data2 is nil")
	} else {
		fmt.Println("data2 is not nil")
	}

	var g Guiwoo
	var gg *Guiwoo

	fmt.Println(g)
	if gg == nil {
		fmt.Println("gg is nil")
	} else {
		fmt.Println("gg is not nil")
	}
	/**
	참조타입(slice,map,channel) 들은 var 를 이용해서 선언하면 nil 로 초기화 된다. nil 이라고해서 메모리가 할당되지 않은것은 아니다.
	반면 사용자가 선언한 구조체 는 값타입이다. 참조타입이 아니라 값타입이기 때문에 nil 로 초기화 되지 않는다.
	반면 *struct{} 는 참조타입이다. 따라서 nil 로 초기화 된다. 포인터는 참조타입의 한종류 이다. 따라서 포인터는 nil 로 초기화 된다.
	*/
}

func inspectSlice(slice []int) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %v\n", i, &slice[i], slice[i])
	}
}

func Test_Slice_OverCapacity(t *testing.T) {
	s := make([]int, 4, 8)
	s3 := s[0:8:8]

	s3 = append(s3, 100)
	fmt.Printf("s : %v, len : %d, cap : %d\n", s, len(s), cap(s))
	inspectSlice(s)
	fmt.Printf("s3 : %v, len : %d, cap : %d\n", s3, len(s3), cap(s3))
	inspectSlice(s3)

	s3[0] = 200

	fmt.Println(s, s3)
}

func TestMap(t *testing.T) {
	m := make(map[string]Guiwoo)

	f, i := m["guiwoo"]
	fmt.Println(f, i)
}
