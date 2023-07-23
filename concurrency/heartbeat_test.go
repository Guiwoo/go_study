package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/**
하트비트
일정 시간간격으로 발생하는 하트 비트 , 하나의 작업단위를 ㅓㅊ라히가 위해 다른일이 얼어날때까지 댁ㅣ중인 동시코드에 유용하다.
고루틴은 이 일이 언제 시작될지 알지못하기 때문에 뭔가 일어나기를 기다리는 동안 잠시 빈둥거리고 있을 것이다.
하트비트는 모든 것이 잘되고 있으며, 그 침묵이 예상된 것이라는 사실을 리스너에게 알리는 방법이다.
*/

func Test_HeartBeat(t *testing.T) {
	doWork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartBeat := make(chan interface{})
		result := make(chan time.Time)

		go func() {
			defer close(heartBeat)
			defer close(result)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartBeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case result <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartBeat, result
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second

	heartbeat, rs := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-rs:
			if ok == false {
				return
			}
			fmt.Printf("result %+v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}

func TestHeartBeat_Panic(t *testing.T) {
	doWork := func(done <-chan interface{}, tick time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		result := make(chan time.Time)

		go func() {
			pulse := time.Tick(tick)
			workGen := time.Tick(2 * tick)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case result <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, result
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second
	heartbeat, rs := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-rs:
			if ok == false {
				return
			}
			fmt.Printf("result %+v", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker go routine is not healthy")
			return
		}
	}
}

func Test_HeartBeat02(t *testing.T) {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartBeatStream := make(chan interface{}, 1)
		workStream := make(chan int)

		go func() {
			defer close(heartBeatStream)
			defer close(workStream)

			for i := 0; i < 10; i++ {
				select {
				case heartBeatStream <- struct{}{}:
				default:
				}
				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}

		}()
		return heartBeatStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			} else {
				fmt.Printf("result is %d\n", r)
			}
		}
	}
}
func DokWorkVer1(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{})
	result := make(chan int)
	go func() {
		defer close(result)
		defer close(heartbeat)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case result <- n:
			}
		}
	}()

	return heartbeat, result
}
func Test_HeartBeat_Writing(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 2, 3, 4, 5}
	_, result := DokWorkVer1(done, intSlice...)
	for i, expected := range intSlice {
		select {
		case r := <-result:
			if r != expected {
				t.Errorf("index %d : expected %d but received %d", i, expected, r)
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("test time out")
		}
	}
}

func Test_HeartBeat_for_delay(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 2, 3, 4, 5, 6}
	heartbeat, results := DokWorkVer1(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; expected != r {
			t.Errorf("index %d expected %d but got %d", i, expected, r)
		}
		i++
	}
}

func TestHeartBeat04(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	heartbeat, results := DokWorkVer1(done, intSlice...)
	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; expected != r {
			t.Errorf("index %d expected %d got %d\n", i, expected, r)
		}
		i++
	}
}

func DoWorkVer2(done <-chan interface{}, tick time.Duration, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(tick)

	numLoop:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					continue numLoop
				}
			}
		}

	}()

	return heartbeat, intStream
}

func TestHeartbeatAllNums(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	const timeout = 2 * time.Second
	hearbeat, results := DoWorkVer2(done, timeout/2, intSlice...)
	<-hearbeat

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf("expected %d got %d", expected, r)
			}
			i++
		case <-hearbeat:
		case <-time.After(timeout):
		}
	}
}

func TestHeartBeatCopyRequest(t *testing.T) {
	doWork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {
		started := time.Now()
		defer wg.Done()

		simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("took %v %v \n", took, id)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg *sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go doWork(done, i, wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("All go answered from %v\n", firstReturned)
}
