package basic

import (
	"fmt"
	"os"
	"path"
)

func ex01() {
	// Make it Blue

	color := "green"
	color = "blue"
	fmt.Println(color)
}

func ex02() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW

	color += "dark"

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	fmt.Println(color)
}
func ex03() {
	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	n := 0.

	// ADD YOUR CODE BELOW

	n = 3.14 * 2

	fmt.Println(n)
}

func ex04() {
	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter     int
		width, height = 5, 6
	)

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT

	// ADD YOUR CODE BELOW
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)
}

func ex05() {
	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	// ADD YOUR CODE BELOW
	lang, version = "go", 2
	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)
}

func ex06() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, isTrue, temp = "Mars", true, 10.5
	fmt.Printf("Airt is Good on %s\nIt's %v\nIt is %v degrees\n", planet, isTrue, temp)
}

func ex07() {
	multi := func() (int, int) { return 5, 4 }
	// ADD YOUR DECLARATIONS HERE
	_, b := multi()
	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)
}

func ex08() {
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"
	fmt.Println(color, color2)
}

func ex09() {
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	// ?
	red, blue = blue, red
	fmt.Println(red, blue)
}

func ex10() {
	// UNCOMMENT THE CODE BELOW:

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

func ex11() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

func ex12() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

func ex13() {
	fmt.Println(float64(5.5))
}

func ex14() {
	age := 2
	fmt.Println(float64(7.5) + float64(age))
}
func ex15() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}
func ex16() {
	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
}
func ex17() {
	fmt.Println("go build -o guiwoo The path is => ", os.Args[0])
	fmt.Println("Got your name is => ", os.Args[1])

	if len(os.Args) < 4 {
		panic("The Argument needs at least more than 3")
	}
	fmt.Println(os.Args[1], os.Args[2], os.Args[3])
}
