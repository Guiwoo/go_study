package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
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
