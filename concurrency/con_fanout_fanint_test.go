package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_fan_01(t *testing.T) {
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
	ran := func() interface{} { return rand.Intn(50000000) }
	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := n - 1; i > 1; i-- {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan int {
		primeStream := make(chan int)
		go func() {
			defer close(primeStream)
			for v := range intStream {
				select {
				case <-done:
					return
				default:
					if isPrime(v) {
						primeStream <- v
					}
				}
			}
		}()
		return primeStream
	}

	take := func(done <-chan interface{}, intStream <-chan int, num int) <-chan int {
		takeStream := make(chan int)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-intStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, ran))

	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))

}
