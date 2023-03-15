package basic

import (
	"fmt"
	"os"
	"strings"
)

// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func i01(age int) {
	// Change this accordingly to produce the expected outputs
	a := true
	switch a {
	case age >= 60:
		fmt.Println("Getting older")
	case age >= 30:
		fmt.Println("Getting wiser")
	case age >= 20:
		fmt.Println("Adulthood")
	case age >= 10:
		fmt.Println("Young blood")
	default:
		fmt.Println("Booting up")
	}
}

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func i02() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool

	if radius >= 200 {
		big = true
	}

	if big != true && isSphere {
		fmt.Println("I don't know.")
	} else {
		fmt.Println("It's a big sphere.")
	}
}

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func i03() {
	os.Stdout.WriteString("GIve an args more than one\n")

	a := os.Args

	fmt.Printf("There are %d : %s\n", len(a)-1, strings.Join(a[1:], " "))
}

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func i04() {
	a := os.Args[1]
	if len(a) > 1 {
		fmt.Println("Give me a Character")
		return
	}

	vowel := []rune{'a', 'e', 'i', 'o', 'u'}
	both := []rune{'w', 'y'}

	for _, v := range a {
		for _, vv := range vowel {
			if vv == v {
				fmt.Printf("%v is a vowel", string(vv))
				return
			}
		}
		for _, b := range both {
			if v == b {
				fmt.Printf("%v is sometimes a vowel, sometimes not", string(b))
				return
			}
		}
	}
	fmt.Printf("%v is a consonant .", a)
}

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func i05() {
}
