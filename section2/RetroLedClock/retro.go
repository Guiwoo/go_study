package RetroLedClock

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func generate() [][]string {
	num := make([][]string, 10)
	for i := range num {
		num[i] = make([]string, 5)
	}

	num[0] = []string{"███", "█ █", "█ █", "█ █", "███"}
	num[1] = []string{"██ ", " █ ", " █ ", " █ ", "███"}
	num[2] = []string{"███", "  █", "███", "█  ", "███"}
	num[3] = []string{"███", "  █", "███", "  █", "███"}
	num[4] = []string{"█ █", "█ █", "███", "  █", "  █"}
	num[5] = []string{"███", "█  ", "███", "  █", "███"}
	num[6] = []string{"███", "█  ", "███", "█ █", "███"}
	num[7] = []string{"███", "  █", "  █", "  █", "  █"}
	num[8] = []string{"███", "█ █", "███", "█ █", "███"}
	num[9] = []string{"███", "█ █", "███", "  █", "███"}
	return num
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func setTime() []int {
	now := time.Now()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	hour1 := hour / 10
	hour2 := hour % 10
	minute1 := minute / 10
	minute2 := minute % 10
	second1 := second / 10
	second2 := second % 10

	return []int{hour1, hour2, -1, minute1, minute2, -1, second1, second2}
}

func drawScreen(arr [][]string) {
	theTime := setTime()

	t := make([][]string, 8)
	for i := range t {
		tx := theTime[i]
		if tx == -1 {
			t[i] = []string{"   ", " █ ", "   ", " █ ", "   "}
			continue
		}
		target := arr[tx]
		t[i] = target
	}
	sec := time.Now().Second()
	for i := 0; i < 5; i++ {
		for j := 0; j < 8; j++ {
			if sec%2 == 0 && (j == 2 || j == 5) {
				fmt.Print("     ")
				continue
			}
			fmt.Print(t[j][i] + "  ")
		}
		fmt.Println()
	}
}

func Start() {
	arr := generate()
	ctx := context.Background()
	clearScreen()
	for {
		c := time.NewTicker(time.Second)
		select {
		case <-c.C:
			fmt.Print("\033[H\033[2J")
			drawScreen(arr)
		case <-ctx.Done():
			fmt.Println("Program is done")
			return
		}
	}
}
