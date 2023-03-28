package main

import "fmt"

func ex01() {
	const (
		winter = 4
		summer = 3
	)
	var books [4]string
	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "The Go is Conqueror the world"
	books[3] = "Hoit"

	fmt.Println("books : %T\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]
	for i, _ := range sBooks {
		sBooks[i] = books[i+1]
	}

	// arr literal
	arr := [3]string{"a", "b"}
	fmt.Println(arr)
}
