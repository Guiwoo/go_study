package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"runtime"
	"sync"
	"text/tabwriter"
	"time"
)

func fork() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func fork2() {
	salutation := "Hello"
	run := make(chan bool, 1)
	go func() {
		run <- true
		salutation = "Welcome"
	}()
	<-run
	fmt.Println(salutation)
}

func fork3() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"Hello", "Greetings", "Good Day"} {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(salutation)
	}
	wg.Wait()
}

func waitGroup() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("%d : Hello\n", id)
	}

	const numGreet = 5
	var wg sync.WaitGroup
	wg.Add(numGreet)
	for i := 0; i < numGreet; i++ {
		go hello(&wg, i)
	}
	wg.Wait()
}

func mutx() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing : %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing : %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

func rwmtx() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)

		beginTestTime := time.Now()
		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Printf("Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
		test(count, &m, m.RLocker())
		test(count, &m, &m)
	}
}

func cond() {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	for true {
		c.Wait()
	}
	c.L.Unlock()
}

func cond2() {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(10 * time.Second)
		c.L.Unlock()
	}
}

func cond3() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		wg.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing windows.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}

func on() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func pool() {
	myPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()
	a := myPool.Get()
	myPool.Put(a)
	myPool.Get()
}

func pool2() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func main() {
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

	time.AfterFunc(2*time.Second, func() {
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

	strat := time.Now()

	things := make([]<-chan interface{}, len(dons))
	for i := 0; i < len(dons); i++ {
		var or int
		if i == 0 {
			or = 0
		} else {
			or = i - 1
		}
		things[i] = orDone(dons[or], repeat(dons[0], i))
	}

	donedone := make(chan interface{})
	defer close(donedone)
	for v := range fanin(donedone, things...) {
		if (time.Since(strat) > 3*time.Second) {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("all done", time.Since(strat))
}
