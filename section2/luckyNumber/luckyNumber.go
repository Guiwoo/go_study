package luckyNumber

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ex01() {
	var (
		a      int
		answer int
	)

	fmt.Println("Lucky Number Game")
	reader := bufio.NewReader(os.Stdin)
	stack := 5
	answer = rand.Intn(101)
	fmt.Println(answer)
	for stack > 0 {
		if _, err := fmt.Fscanln(reader, &a); err != nil {
			panic("Failure to parse number")
		}
		switch {
		case a == answer:
			fmt.Printf("You won the left stack is %d\n", stack)
			break
		case a > answer:
			fmt.Printf("Wrong Number should go lower left stack is : %d\n", stack)
		default:
			fmt.Printf("Wrong Number should go higher left stack left stack is :%d\n", stack)
		}
		stack--

	}
	fmt.Printf("The answer is %d", answer)
}
