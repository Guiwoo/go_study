package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})

	randStream := newRandStream(done)
	fmt.Println("3 random ints :")
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d : %d\n", i, <-randStream)
	}
	close(done)

	time.Sleep(1 * time.Second)
}

func Test_Pattern_07(t *testing.T) {
	// 1개 이상의 채널들을 보낼수 있다 <- 보내는 애들로만
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

func Test_Pattern_08(t *testing.T) {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		response := make(chan *http.Response)
		go func() {
			defer close(response)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
					continue
				}
				select {
				case <-done:
					return
				case response <- resp:
				}
			}
		}()
		return response
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://www.naver.com", "http://localhost:8080"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response : %v\n", response.Status)
	}
}

func Test_Pattern_09(t *testing.T) {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				rs := Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- rs:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://www.naver.com", "http://localhost:8080"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error : %v\n", result.Error)
			continue
		}
		fmt.Printf("Response : %v\n", result.Response.Status)
	}

	fmt.Println()

	urls = []string{"https://www.google.com", "https://www.naver.com", "http://localhost:8080", "a", "b", "c", "d", "e", "f", "g"}
	errCount := 0
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error : %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response : %v\n", result.Response.Status)
	}

}

func Test_Pattern_10(t *testing.T) {
	//Reader 파이프
	reader := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, reader(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}

func Test_Pattern_11(t *testing.T) {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, repeat int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < repeat; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	rr := func() interface{} { return rand.Int() }

	for v := range take(done, repeatFn(done, rr), 10) {
		fmt.Println(v)
	}
}

func Test_Pattern_12(t *testing.T) {
	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	var message string
	for v := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += v
	}
	fmt.Printf("message : %s...", message)
}

func Benchmark_Pattern_01(b *testing.B) {
	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	for range toString(done, take(done, repeat(done, "I", "am."), 1000000)) {
	}
}

func Benchmark_Pattern_02(b *testing.B) {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	for range take(done, repeat(done, "I", "am."), 1000000) {
	}
}
