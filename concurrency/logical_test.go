package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestLogicalProcessor(t *testing.T) {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for cnt := 0; cnt < 200; cnt++ {
			for char := 'a'; char <= 'z'; char++ {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()
		for cnt := 0; cnt < 200; cnt++ {
			for char := 'A'; char <= 'Z'; char++ {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}()
	time.Sleep(10 * time.Second)
	wg.Wait()
}
