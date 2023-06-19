package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
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

func Test_fan_02(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

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

	isPrime := func(n interface{}) bool {
		x := n.(int)
		if x < 2 {
			return false
		}
		for i := x - 1; i > 1; i-- {
			if x%i == 0 {
				return false
			}
		}
		return true
	}

	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
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

	take := func(done <-chan interface{}, stream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-stream:
				}
			}
		}()
		return takeStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func Test_fan_03(t *testing.T) {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			defer fmt.Println("orDone 함수 종료")
			for {
				select {
				case <-done:
					fmt.Println("Get Value out from case v,ok := <-c")
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
						fmt.Println("Get the value")
					case <-done:
						fmt.Println("get value from done")
					}
				}
			}
		}()
		return valStream
	}

	testing := func(done <-chan interface{}) <-chan interface{} {
		testingStream := make(chan interface{})
		go func() {
			defer close(testingStream)
			defer fmt.Println("testing 함수종료")
			for {
				select {
				case <-done:
					return
				case testingStream <- struct{}{}:
					fmt.Println("Send the value")
				}
				time.Sleep(1 * time.Second)
			}
		}()
		return testingStream
	}

	done := make(chan interface{})
	done2 := make(chan interface{})

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("close done")
		defer close(done)
	})
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("close done2")
		defer close(done2)
	})

	for a := range orDone(done2, testing(done)) {
		fmt.Println(a)
	}
}

func Test_fan_04(t *testing.T) {
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
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out2)
			defer close(out1)

			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, orDone(done, take(done, repeat(done, 1, 2, 3), 10)))
	for a := range out1 {
		fmt.Printf("out1 : %v, out2 : %v\n ", a, <-out2)
	}
}

func Test_fan_05(t *testing.T) {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valueStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valueStream
	}

	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valueStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valueStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}

func Test_CPU_CHECK(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}

func Test_OR_Done(t *testing.T) {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer fmt.Println("repeat closed ", values)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
				time.Sleep(1 * time.Second)
			}
		}()
		return valueStream
	}
	orDone := func(done, value <-chan interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			defer fmt.Println("Close value stream")
			for {
				select {
				case <-done:
					fmt.Println("Close done in first select")
					return
				case v, ok := <-value:
					if ok == false {
						fmt.Println("Closed value channel")
						return
					}
					select {
					case valueStream <- v:
						fmt.Println("Value sent")
					case <-done:
						fmt.Println("Close done in second select")
					}
				}
			}
		}()
		return valueStream
	}

	dons := make([]chan interface{}, runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		dons[i] = make(chan interface{})
	}

	time.AfterFunc(1*time.Second, func() {
		close(dons[2])
	})
	time.AfterFunc(5*time.Second, func() {
		for i := 0; i < runtime.NumCPU(); i++ {
			if i == 2 {
				continue
			}
			close(dons[i])
		}
		defer fmt.Println("all dons closed")
	})

	fanin := func(done <-chan interface{}, chans ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		valueStream := make(chan interface{})

		output := func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}

		wg.Add(len(chans))
		for _, c := range chans {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(valueStream)
		}()

		return valueStream
	}

	things := make([]<-chan interface{}, len(dons))
	for i := 0; i < len(dons); i++ {
		var or int
		if i == 0 {
			or = 0
		} else {
			or = i - 1
		}
		things[i] = orDone(dons[or], repeat(dons[i], i))
	}

	donedone := make(chan interface{})
	defer close(donedone)
	for v := range fanin(donedone, things...) {
		fmt.Println(v)
	}

	fmt.Println("all done")
}

func Test_case_done_why_need_it_fuck(t *testing.T) {
	tt := func() <-chan interface{} {
		val := make(chan interface{})
		time.AfterFunc(10*time.Second, func() {
			close(val)
		})
		go func() {
			defer close(val)
			for {
				val <- "sibal"
				time.Sleep(1 * time.Second)
			}
		}()
		return val
	}

	mid := func(done, c <-chan interface{}) <-chan interface{} {
		val := make(chan interface{})
		go func() {
			defer close(val)
			for {
				select {
				case val <- c:
					fmt.Println("값 전달")
				case <-done:
					fmt.Println("5초 기다려야해서 다른거 프린트")
				}
			}
		}()
		return val
	}

	done := make(chan interface{})
	var i int
	for a := range mid(done, tt()) {
		i++
		if i == 5 {
			break
		}
		fmt.Println(a)
	}
}
