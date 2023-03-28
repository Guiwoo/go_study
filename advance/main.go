package main

import (
	"fmt"
	"sync"
	"time"
)

type Command struct {
	name string
	done chan bool
}

func main() {
	var wg sync.WaitGroup
	commandCh := make(chan Command)
	goroutines := make(map[string]chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		doneCh := make(chan bool)
		goroutines[fmt.Sprintf("goroutine-%d", i)] = doneCh
		go func(id int, doneCh chan bool) {
			defer wg.Done()
			name := fmt.Sprintf("goroutine-%d", id)
			tick := time.Tick(time.Second * 10)
			for {
				select {
				case <-tick:
					fmt.Printf("%s is running \n", name)
				case <-doneCh:
					fmt.Printf("%s is terminated \n", name)
					fmt.Println()
					fmt.Println("✅ Enter the name of the goroutine to kill:")
					return
				case cmd := <-commandCh:
					if cmd.name == name {
						cmd.done <- true
						doneCh <- true
					}
				}
			}
		}(i, doneCh)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("✅ Enter the name of the goroutine to kill:")
		var name string
		_, err := fmt.Scanln(&name)
		if err != nil {
			return
		}
		if doneCh, ok := goroutines[name]; ok {
			commandCh <- Command{name: name, done: doneCh}
			<-doneCh
		} else {
			fmt.Println("Invalid goroutine name. Please try again.")
			i--
		}
	}
	wg.Wait()
	fmt.Println("All goroutine killed")
}
