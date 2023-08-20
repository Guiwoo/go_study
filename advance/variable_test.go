package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_Variable_String_Rune(t *testing.T) {
	a := "안녕"

	for _, v := range []rune(a) {
		fmt.Printf("%c\n", v)
	}

	for _, v := range a {
		fmt.Printf("%c\n", v)
	}
}

func Test_Memeory_Address(t *testing.T) {
	type ex struct {
		counter int64
		pi      float32
		flag    bool
	}

	type ex2 struct {
		flag    bool
		counter int64
		pi      float32
	}

	var e ex
	var e2 ex2

	fmt.Println(unsafe.Sizeof(e), unsafe.Sizeof(e2))
}

type user struct {
	name, email string
}

func stayOnStack() user {
	u := user{
		name:  "Ho",
		email: "email",
	}
	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "Ho",
		email: "email",
	}
	return &u
}

func Test_Pointer_Address(t *testing.T) {
	fmt.Println(stayOnStack())
	fmt.Println(escapeToHeap())
	// go test -gcflags '-m -l' advance/variable_test.go
}

/*
*
가비지 컬렉션
*/
func Test_Garbage_Collect(t *testing.T) {

}
