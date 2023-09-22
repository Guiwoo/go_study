package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNatsClient(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		n := NewNatsClient()
		n.subscribe("hoit", func(data string) error {
			time.Sleep(2 * time.Second)
			fmt.Println("함수 1번째", data)
			return nil
		})
	}()

	go func() {
		n := NewNatsClient()
		n.subscribe("hoit", func(data string) error {
			fmt.Println("함수 2번째", data)
			return nil
		})
	}()

	wg.Wait()
}
