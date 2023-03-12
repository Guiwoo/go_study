package main

import (
	"fmt"
	"os"
	"path"
)

func ex02() {
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println(file)
}

// When use a short declaration ?
// if you don't know initial value ? do not use it
// when you need a package scoped variable

var globalValue int

func ex03() {
	// score := 0 // DONT
	var score int // 👍Already score = 0

	// better readability grouping the value like  this
	var (
		video    string
		duration int
		current  int
	)
	fmt.Println(score, video, duration, current)
}

// But Go Developer love short init way
func ex04() {
	// If you know the specific value ? use this
	width, height := 100, 160 //👍 Good

	width, color := 50, "red" // 👍Look better

	fmt.Println(width, height, color)
}

// type conversion
func ex05() {
	// You can't use value belong to different types together
	speed := 100 // int
	force := 2.5 // float64

	// You need to explicitly convert the values
	// conversion can change value it self
	speed = speed * int(force)

	speed = int(float64(speed) * force)
}

func ex06() {
	// A slice can store multiple values
	var Args []string // Args's type is slice of strings
	// Args[0] => path to the program Path to the program
	Args = os.Args
	fmt.Printf("%#v\n", Args)
	fmt.Println("Paht", Args[0])
	fmt.Println("1st", Args[1])
	fmt.Println("Number of items", len(Args))
}

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}
