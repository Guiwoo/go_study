package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func lowercase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
	fmt.Println()
}

func uppercase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
	fmt.Println()
}

func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s :%d \n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func timeSlicing() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Go routine start")
	go func() {
		defer wg.Done()
		lowercase()
	}()

	go func() {
		defer wg.Done()
		uppercase()
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

var counter int64

func raceCondition() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		go func() {
			defer wg.Done()
			for count := 0; count < 2; count++ {
				value := counter
				runtime.Gosched()
				value++
				counter = value
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter : ", counter)
}
func atomicWay() {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		go func() {
			defer wg.Done()
			for cnt := 0; cnt < 2; cnt++ {
				atomic.AddInt64(&counter, 1)
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter : ", counter)
}

func mutexWay(mutex *sync.Mutex) {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			defer wg.Done()
			for count := 0; count < 2; count++ {

				mutex.Lock()
				{
					value := counter
					value++
					counter = value
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println("Final Counter : ", counter)
}

var (
	data      []string
	rwMutex   sync.RWMutex
	readCount int64
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	writer := func(i int) {
		rwMutex.Lock()
		defer rwMutex.Unlock()
		{
			rc := atomic.LoadInt64(&readCount)
			fmt.Printf("*****> : Performing Write : RCount[%d]\n", rc)
			data = append(data, fmt.Sprintf("String : %d", i))
		}
	}

	reader := func(i int) {
		rwMutex.RLock()
		defer rwMutex.RUnlock()
		{
			rc := atomic.AddInt64(&readCount, 1)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			fmt.Printf("%d : Performing Read : :Length[%d] RCount[%d]\n", i, len(data), rc)
			atomic.AddInt64(&readCount, -1)
		}
	}

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
	}()

	for i := 0; i < 8; i++ {
		go func(i int) {
			for {
				reader(i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("")
}
