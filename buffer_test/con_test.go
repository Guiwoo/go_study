package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func workerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for j := range jobs {
		result <- process(j)
	}
}

func process(job int) int {
	time.Sleep(1 * time.Second)
	return job * 2
}

func TestWorkerPool(t *testing.T) {
	numJobs := 100
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	workerPool(5, jobs, results)

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for v := range results {
		fmt.Println(v)
	}

	fmt.Println("Done")
}

func worker2(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			output <- process(n)
		}
	}()
	return output
}

func fanOut(input <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		channels[i] = worker2(input)
	}
	return channels
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	stream := make(chan int)

	p := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			stream <- i
		}
	}
	wg.Add(len(channels))

	for _, c := range channels {
		go p(c)
	}

	go func() {
		wg.Wait()
		close(stream)
	}()

	return stream
}

func TestFanInFanOut(t *testing.T) {
	input := make(chan int, 100)
	worker3 := fanOut(input, 5)
	results := fanIn(worker3...)

	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
	}()

	for r := range results {
		fmt.Printf("%+v\n", r)
	}
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func print(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func TestPipeLine(t *testing.T) {
	//3중 포문을 pipeline 화 시켜서 해보자
	print(double(square(generator(1, 2, 3, 4, 5))))
}

func worker4(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker 4 is done !")
		default:
			fmt.Println("Worker 4 is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go worker4(ctx, i)
	}

	time.Sleep(time.Second * 5)

	fmt.Println("Main is Done")
}
