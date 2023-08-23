package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func testStream() <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		t1 := time.NewTicker(1 * time.Second)
		t10 := time.Tick(10 * time.Second)
		defer close(stream)
		for {
			select {
			case <-t1.C:
				stream <- "thing"
			case <-t10:
				return
			}
		}
	}()
	return stream
}

func Test_go_profile_do_not_take_value(t *testing.T) {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	for ok := range testStream() {
		fmt.Println(ok)
	}
}
