package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//제한

func Test_Pattern_01(t *testing.T) {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func Test_Pattern_02(t *testing.T) {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(rs <-chan int) {
		for r := range rs {
			fmt.Println(r)
		}
		fmt.Println("Done")
	}

	rs := chanOwner()
	consumer(rs)
}

func Test_Pattern_03(t *testing.T) {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff []byte
		for _, b := range data {
			buff = append(buff, b)
		}
		fmt.Println(string(buff))
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func Test_Pattern_04(t *testing.T) {
	dowork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer close(completed)
			defer fmt.Println("Done")
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	dowork(nil)
	fmt.Println("Done.")
}

func Test_Pattern_05(t *testing.T) {
	dowork := func(done <-chan bool, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited")
			//defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Printf("Received : %s\n", s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan bool)
	terminated := dowork(done, nil)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

func Test_Pattern_06(t *testing.T) {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints :")
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d : %d\n", i, <-randStream)
	}
}