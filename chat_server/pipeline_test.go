package main

import (
	"sync"
	"testing"
	"time"
)

func Test_PipeLine(t *testing.T) {
	repeat := func(done <-chan interface{}, i string) <-chan interface{} {
		value := make(chan interface{})
		go func() {
			defer close(value)
			for {
				select {
				case <-done:
					return
				case value <- i:
					time.Sleep(1 * time.Second)
				}
			}
		}()
		return value
	}

	fanOut := make([]<-chan interface{}, 8)
	for i := 0; i < len(fanOut); i++ {
		fanOut[i] = repeat(nil, "Hello")
	}

	fanIn := func(done <-chan interface{}, chans ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multipleStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multipleStream <- i:
				}
			}
		}

		wg.Add(len(chans))
		for _, c := range chans {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multipleStream)
		}()

		return multipleStream
	}

	done := make(chan interface{})

	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	for v := range fanIn(done, fanOut...) {
		t.Log(v)
	}
}
