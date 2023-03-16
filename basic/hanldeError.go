package basic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func err01() {
	a := os.Args[1]

	if _, err := strconv.Atoi(a); err != nil {
		errors.New(fmt.Sprintf("%s is not a number", a))
	}
	fmt.Println("This is a number")
}

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func errorEx02() {
	var (
		n   int
		err error
	)
	if a := os.Args; len(a) != 2 {
		fmt.Println("give me a number")
	} else {
		if n, err = strconv.Atoi(a[1]); err != nil {
			fmt.Println("Cna't convert to integer")
		}
	}

	if n > 17 {
		fmt.Println("R-Rated")
	} else if n > 13 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("PG-Rated")
	}
}

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//
//	Handle the error. If the number is not a valid number,
//	or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//
//	go run main.go 16
//	  16 is an even number and it's divisible by 8
//
//	go run main.go 4
//	  4 is an even number
//
//	go run main.go 3
//	  3 is an odd number
//
//	go run main.go
//	  Pick a number
//
//	go run main.go ABC
//	  "ABC" is not a number
//
// ---------------------------------------------------------
func errorEx03() {
	var a []string
	if a = os.Args; len(a) != 2 {
		fmt.Printf("Give me  a number")
	}

	n, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Printf("%s is not a number", a[1])
	}

	if n%2 == 0 {
		fmt.Printf("%d is even number", n)
	} else {
		fmt.Printf("%d is odd number", n)
	}
}

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//	Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//
//	go run main.go
//	  Give me a year number
//
//	go run main.go eighties
//	  "eighties" is not a valid year.
//
//	go run main.go 2018
//	  2018 is not a leap year.
//
//	go run main.go 2100
//	  2100 is not a leap year.
//
//	go run main.go 2019
//	  2019 is not a leap year.
//
//	go run main.go 2020
//	  2020 is a leap year.
//
//	go run main.go 2024
//	  2024 is a leap year.
//
// ---------------------------------------------------------
func errorEx04() {
	a := os.Args[1]

	n, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Give me a number ")
	} else if n <= 0 {
		fmt.Println("Pleas give me a number")
	}

	if n%4 == 0 && n%100 == 0 {
		fmt.Println("This is leaf year", n)
	}
}

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func errEx06() {
	var (
		a   []string
		n   int
		err error
	)

	if a = os.Args; len(a) != 2 {
		fmt.Println("give me a number")
		return
	}

	if n, err = strconv.Atoi(a); err != nil {
		fmt.Println("Give me a number ")
	} else if n <= 0 {
		fmt.Println("Pleas give me a number")
	} else if n%4 == 0 && n%100 == 0 {
		fmt.Println("This is leaf year", n)
	}
}

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func errorEx08() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days := 28

	month := os.Args[1]

	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
