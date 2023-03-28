package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Switch
/**
Switch's condition determines the type of the case conditions
*/

func swEx01() {
	city := "Tokyo"
	// Each Condition case specific following type
	switch city {
	case "Paris":
		fmt.Println("France")
		// Not necessary to use break case
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Somewhere in the Earth")
	}

	//FallThrough statement work
	a := os.Args[1]
	n, _ := strconv.Atoi(a)

	switch {
	case n > 100:
		fmt.Print("Big ")
		fallthrough
	case n > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("Number\n")
	}

	//Short switch
	switch i := 10; true {
	case i > 0:
		fmt.Println("Something")
	default:
		fmt.Println("Zero")
	}
	t := time.Now().Hour()
	//6 ~ 11
	//12 ~ 18
	// 18 ~ 22
	// 23 ~ 5

	switch {
	case 6 < t && t < 11:
		fmt.Println("Good morning")
	case 12 < t && t < 18:
		fmt.Println("Good afternoon")
	case 18 < t && t < 22:
		fmt.Println("Good evening")
	default:
		fmt.Println("Mid Night")
	}
}

/**
So when do i need to use if vs switch ?
*/

func main() {

}
