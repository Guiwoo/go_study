package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var sem = make(chan int, 10)
	process := func(r *http.Request) {
		fmt.Println(r.Method)
	}

	handler := func(r *http.Request) {
		sem <- 1
		process(r)
		<-sem
	}

	Serve := func(queue <-chan *http.Request) {
		for {
			req := <-queue
			go handler(req)
		}
	}

	repeat := func(done <-chan interface{}) <-chan *http.Request {
		t := time.NewTicker(1 * time.Second)
		stream := make(chan *http.Request)
		go func() {
			defer close(stream)
			for {
				select {
				case <-done:
					return
				case <-t.C:
					stream <- &http.Request{Method: fmt.Sprintf("request %+v", <-t.C)}
				}
			}
		}()
		return stream
	}

	done := make(chan interface{})

	a := repeat(done)
	Serve(a)

	go func() {
		defer close(done)

		time.Sleep(30 * time.Second)
		done <- ""
	}()
	fmt.Println("done")
}

type Vector []float64

func (v Vector) Op(f float64) float64 {
	return 0
}
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
}

const numCpu = 4

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCpu)
	for i := 0; i < numCpu; i++ {
		go v.DoSome(i*len(v)/numCpu, i+1*len(v)/numCpu, u, c)
	}

	for i := 0; i < numCpu; i++ {
		<-c
	}
	// 모두종료
}

func Test022(t *testing.T) {
	timerChan := make(chan time.Time)
	go func() {
		time.Sleep(20 * time.Second)
		timerChan <- time.Now()
	}()

	fmt.Println(<-timerChan)
}

func Test023(t *testing.T) {
	bfChan := make(chan interface{}, 20)

	go func() {
		defer close(bfChan)

		for i := 0; i < 50; i++ {
			fmt.Println("a 를 넣습니다. ", i+1)
			bfChan <- "a" + fmt.Sprintf("%d", i+1)
		}

	}()
	for rst := range bfChan {
		fmt.Println("a 를 빼고 있습니다.", rst)
		time.Sleep(time.Second * 1)
	}
}

type Work struct {
	x, y, z int
}

func sleep(a int) {
	time.Sleep(time.Second * time.Duration(a))
}
func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		sleep(w.z)
		out <- w
	}
}

func TestLoadBalancer(t *testing.T) {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 20; i++ {
		go worker(in, out)
	}
	//go sendLostOfWork(in)
	//receiveLostOfResuls(out)
}
