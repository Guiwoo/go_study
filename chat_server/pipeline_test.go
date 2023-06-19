package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_PipeLine(t *testing.T) {
	repeat := func(done <-chan interface{}) <-chan interface{} {
		value := make(chan interface{})
		repeatDone := make(chan interface{})
		go func() {
			defer close(value)
			defer close(repeatDone)
			for {
				select {
				case <-done:
					return
				case value <- "after 1sec ":
				}
				time.Sleep(1 * time.Second)
			}
		}()
		go func() {
			<-done
			close(repeatDone)
		}()
		return value
	}
	chatSignal := func(done, value <-chan interface{}) <-chan interface{} {
		chatStream := make(chan interface{})
		go func() {
			defer close(chatStream)
			for v := range value {
				select {
				case <-done:
					return
				case chatStream <- v.(string) + "from chatSignal":
				}
			}
		}()
		return chatStream
	}
	enterSignal := func(done, value <-chan interface{}) <-chan interface{} {
		enterStream := make(chan interface{})
		go func() {
			defer close(enterStream)
			for v := range value {
				select {
				case <-done:
					return
				case enterStream <- v.(string) + " from enterSignal":
				}
			}
		}()
		return enterStream
	}
	exitSignal := func(done, value <-chan interface{}) <-chan interface{} {
		exitSignal := make(chan interface{})
		go func() {
			defer close(exitSignal)
			for v := range value {
				select {
				case <-done:
					return
				case exitSignal <- v.(string) + " from exitSignal":
				}
			}
		}()
		return exitSignal
	}

	chain := func(done <-chan interface{}, chans ...<-chan interface{}) <-chan interface{} {
		organize := make(chan interface{})
		chaining := func(cc <-chan interface{}) {
			defer close(organize)
			for {
				select {
				case <-done:
					return
				case v, ok := <-cc:
					if !ok {
						return
					} else {
						organize <- v.(string) + " from chaining \n"
					}
				}
			}
		}
		var wg sync.WaitGroup
		wg.Add(len(chans))
		for _, c := range chans {
			go chaining(c)
		}

		go func() {
			wg.Wait()
			close(organize)
		}()
		return organize
	}

	done := make(chan interface{})
	time.AfterFunc(1*time.Second, func() {
		close(done)
	})

	chat := chatSignal(done, repeat(done))
	enter := enterSignal(done, repeat(done))
	exit := exitSignal(done, repeat(done))

	for v := range chain(done, chat, enter, exit) {
		fmt.Printf("%v\n", v)
	}
}
