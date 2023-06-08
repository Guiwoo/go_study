package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host : %+v", err)
		}
		if _, err = io.ReadAll(conn); err != nil {
			b.Fatalf("cannot read : %+v", err)
		}

		conn.Close()
	}
}

func Test_Mutex(t *testing.T) {
	var count int
	var lock sync.Mutex

	increment := func() {
		defer lock.Unlock()

		lock.Lock()
		count++
		fmt.Println("Added Count current is : ", count)
	}
	decrement := func() {
		defer lock.Unlock()

		lock.Lock()
		count--
		fmt.Println("Sub Count current is : ", count)
	}

	var proc sync.WaitGroup
	for i := 0; i < 5; i++ {
		proc.Add(1)
		go func() {
			defer proc.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		proc.Add(1)
		go func() {
			defer proc.Done()
			decrement()
		}()
	}

	proc.Wait()

	fmt.Println("All Tasks done")
}

func Test_Once(t *testing.T) {
	var count int
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Println(count)
}

func Test_Once2(t *testing.T) {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}

func Test_Pool1(t *testing.T) {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

	m := sync.Map{}
	_, ok := m.LoadOrStore("instance", "something")
	if !ok {
		fmt.Println("Creating new instance")
	}
}

func Test_Began(t *testing.T) {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func Test_BufferedCh(t *testing.T) {
	var strdoutBuff bytes.Buffer
	defer strdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&strdoutBuff, "Producer Done")

		for i := 0; i < 4; i++ {
			fmt.Fprintf(&strdoutBuff, "Sending : %d\n", i)
			intStream <- i
		}
	}()

	for i := range intStream {
		fmt.Fprintf(&strdoutBuff, "Received %v\n", i)
	}
}

func Test_Channel_Owner(t *testing.T) {
	chatOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chatOwner()
	for rs := range resultStream {
		fmt.Println(rs)
	}

	fmt.Println("Done Receiving!")
}

func Test_Select_01(t *testing.T) {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	case c3 <- struct{}{}:
		fmt.Println("c3")
	}
}

func Test_Select_02(t *testing.T) {
	start := time.Now()

	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.", time.Since(start))
	}
}

func Test_Multiple_Channel(t *testing.T) {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count : %d\nc2Count : %d\n", c1Count, c2Count)
}

func Test_Select_03(t *testing.T) {
	start := time.Now()

	var c1, c2 <-chan int

	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Println("In default after ", time.Since(start))
	}
}

func Test_Select_04(t *testing.T) {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
