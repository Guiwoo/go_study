package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
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

func Test_Slice(t *testing.T) {
	data := make([]string, 5, 8)
	data[0] = "apple"
	data[1] = "banana"
	data[2] = "citron"
	data[3] = "durian"
	data[4] = "egg plant"

	slice2 := data[2:4:4]
	slice2 = append(slice2, "add1")
	slice2 = append(slice2, "add2")
	slice2 = append(slice2, "add3")
	slice2 = append(slice2, "add4")
	slice2[5] = "changed"

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("[%d] %p %s \n", i, &slice2[i], slice2[i])
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf("[%d] %p %s \n", i, &data[i], data[i])
	}
}

func Test_SliceReference(t *testing.T) {
	x := make([]int, 7)
	for i := 0; i < len(x); i++ {
		x[i] = i * 100
	}
	twoHundred := &x[1]
	x = append(x, 800)
	x[1]++

	fmt.Println(x[1], *twoHundred)
}

func Test_UTF8(t *testing.T) {
	s := "세계 means world"

	var buf [utf8.UTFMax]byte
	for i, r := range s {
		rl := utf8.RuneLen(r)

		si := i + rl

		copy(buf[:], s[i:si])
		fmt.Printf("%2d: %q: codepoint: %#6x encode bytes : %#v\n", i, r, r, buf)
	}
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s <%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}

func TestMethod(t *testing.T) {
	users := []user{
		{"bill", "bill@email.com"},
		{"hoanh", "hoanh@email.com"},
	}

	for _, u := range users {
		u.changeEmail("changed")
	}
}
