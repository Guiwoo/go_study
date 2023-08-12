package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func Test_Queue_01(t *testing.T) {

	repeat := func(done <-chan interface{}, i int) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- i:
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, n int, val <-chan interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for i := 0; i < n; i++ {
				select {
				case <-done:
					return
				case valueStream <- val:
				}
			}
		}()
		return valueStream
	}

	sleep := func(done <-chan interface{}, t time.Duration, val <-chan interface{}) <-chan interface{} {
		tick := time.NewTicker(t)
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			select {
			case <-done:
				return
			case <-tick.C:
				valueStream <- val
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)

	zeros := take(done, 3, repeat(done, 0))
	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long

	for a := range pipeline {
		fmt.Println(a)
	}
}

func repeat(done <-chan interface{}, target byte) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- target:
			}
		}
	}()
	return valueStream
}

func take(done, val <-chan interface{}, rp int) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for i := 0; i < rp; i++ {
			select {
			case <-done:
				return
			case v := <-val:
				switch t := v.(type) {
				case byte:
					valueStream <- t
				default:
				}
			}
		}
	}()
	return valueStream
}

func tmpFileOrFatal() *os.File {
	file, err := os.CreateTemp("", "tmp")
	if err != nil {
		log.Fatalf("error : %s", err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()

	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}

}

func Benchmark_Queue_02(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}
func Benchmark_Queue_02_01(b *testing.B) {
	buf := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(buf))
}
