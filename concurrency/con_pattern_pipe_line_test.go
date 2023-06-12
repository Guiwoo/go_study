package main

import (
	"fmt"
	"testing"
)

func multiply(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func add(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

type Multi struct {
	value []int
}

func (m *Multi) multiply(multiplier int) *Multi {
	for i := range m.value {
		m.value[i] *= multiplier
	}
	return m
}
func (m *Multi) add(adder int) *Multi {
	for i := range m.value {
		m.value[i] += adder
	}
	return m
}

func Test_pipe_line(t *testing.T) {
	ints := []int{1, 2, 3, 4}

	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}

	fmt.Println()

	m := &Multi{value: ints}
	for _, v := range m.multiply(2).add(1).multiply(2).value {
		fmt.Println(v)
	}
}

func Test_pipe_line_01(t *testing.T) {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addStream <- i + additive:
				}
			}
		}()
		return addStream
	}

	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeLine := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeLine {
		fmt.Println(v)
	}
}

type Cal struct {
	done chan interface{}
	ch   chan int
}

func (c *Cal) generate(ints ...int) *Cal {
	cc := make(chan int)
	go func() {
		defer close(cc)
		for _, i := range ints {
			select {
			case <-c.done:
				return
			case cc <- i:
			}
		}
		c.ch = cc
	}()
	c.ch = cc
	return c
}

func (c *Cal) multiply(multiplier int) *Cal {
	cc := make(chan int)
	go func() {
		defer close(cc)
		for i := range c.ch {
			select {
			case <-c.done:
				return
			case cc <- int(i) * multiplier:
			}
		}
		c.ch = cc
	}()
	c.ch = cc
	return c
}

func (c *Cal) adder(adder int) *Cal {
	cc := make(chan int)
	go func() {
		defer close(cc)
		for i := range c.ch {
			select {
			case <-c.done:
				return
			case cc <- int(i) + adder:
			}
		}
		c.ch = cc
	}()
	c.ch = cc
	return c
}

func Test_pipe_line_test(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	c := &Cal{done: done}
	c = c.generate(1, 2, 3, 4).multiply(2).adder(1).multiply(2)
	for v := range c.ch {
		fmt.Println(v)
	}
}
