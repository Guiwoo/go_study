package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestNatsClient(t *testing.T) {
	n := NewNatsClient()
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		n.subscribe("hoit", func(data string) error {
			fmt.Println(data)
			wg.Done()
			return nil
		})
	}()

	wg.Wait()
}
