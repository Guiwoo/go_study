package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func s01() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println(path)
	path = `c:\program files\duper super\fun.txt\n
c:\program files\really\funny.png\n`
}

func s02() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	fmt.Println(json)
	json = `
{
	Items:[{
		Item:{
			name : Teddy Bear
		}
	}]
}
`
	fmt.Println(json)
}

func s03() {
	// uncomment the code below
	name := os.Args[1]

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi ` + name + `how are you?`

	fmt.Println(msg)
}

func s04() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	length := len([]rune(os.Args[1]))
	fmt.Println(length)
}

func s05() {
	msg := os.Args[1]
	deco := strings.Repeat("!", len(msg))
	s := deco + msg + deco

	fmt.Println(s)
}
func s06() {
	//  1. Look at the documentation of strings package
	//  2. Find a function that changes the letters into lowercase
	//  3. Get a value from the command-line
	//  4. Print the given value in lowercase letters
	//
	// HINT
	//  Check out the strings package from Go online documentation.
	//  You will find the lowercase function there.
	//
	// INPUT
	//  "SHEPARD"
	//
	// EXPECTED OUTPUT
	//  shepard

	fmt.Println(strings.ToUpper(os.Args[1]))
}

func s07() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}

func s08() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	utf8.RuneCountInString(name)
	fmt.Println(len(strings.TrimRight(name, " ")))
}

func p01() {
	fmt.Println("I'm 30 years old.")

	a := "Guiwoo"
	b := "Park"
	fmt.Printf("My name is %s and my lastname is %s\n", a, b)

	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE
	fmt.Printf("These are %t claims\n", tf)

	tmepf := 7.2
	fmt.Printf("Temperature is %.1f cellcius degree", tmepf)

	fmt.Printf("\"Hello World\"")

	c := 3
	fmt.Printf("Type of %d is %[1]T", c)

	d := 3.14
	fmt.Printf("Type of %.2f is %[1]T\n", d)

	e := "Hello"
	fmt.Printf("Type of %s is %[1]T\n", e)

	fmt.Printf("Type of %t is %[1]T\n", true)
}
