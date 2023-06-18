package main

import "testing"

func Test_PipeLine(t *testing.T) {
	repeat := func() <-chan interface{} {
		value := make(chan interface{})
		go func() {
			defer close(value)
			for {

			}
		}()
		return value
	}
}
